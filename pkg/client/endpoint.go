package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/compspec/jgf-go/pkg/jgf"
	pb "github.com/converged-computing/fluxion/pkg/fluxion-grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func (c *FluxionClient) Match(ctx context.Context, in *pb.MatchRequest, opts ...grpc.CallOption) (*pb.MatchResponse, error) {
	response := &pb.MatchResponse{}
	if !c.Connected() {
		return response, errors.New("client is not connected")
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// An error here is an error with making the request
	response, err := c.service.Match(ctx, in)
	if err != nil {
		fmt.Printf("[Match] did not receive any match response: %v\n", err)
		return response, err
	}
	return response, err
}

// Satisfy request for resources
func (c *FluxionClient) Satisfy(ctx context.Context, in *pb.SatisfyRequest, opts ...grpc.CallOption) (*pb.SatisfyResponse, error) {
	response := &pb.SatisfyResponse{}
	if !c.Connected() {
		return response, errors.New("client is not connected")
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// An error here is an error with making the request
	response, err := c.service.Satisfy(ctx, in)
	if err != nil {
		fmt.Printf("[Satisfy] did not receive response: %v\n", err)
		return response, err
	}
	return response, err
}

// PartialCancel takes a set of core ids and a jobid and issues a cancel request
// We have the entire graph here, so we use it to assemble the cancel request.
func (c *FluxionClient) PartialCancel(ctx context.Context, in *pb.PartialCancelRequest, opts ...grpc.CallOption) (*pb.PartialCancelResponse, error) {
	response := pb.PartialCancelResponse{Status: pb.PartialCancelResponse_PARTIAL_CANCEL_ERROR}
	if c.jgf == nil {
		response.Status = pb.PartialCancelResponse_PARTIAL_CANCEL_DISABLED
		return &response, fmt.Errorf("Partial cancel is not enabled.")
	}

	// Prepare a graph for a cancel response
	graph := jgf.NewFluxJGF()
	seenEdges := map[string]bool{}

	// addNewEdges to the graph Edges if we haven't yet
	var addNewEdges = func(path string) {
		addEdges, ok := c.edgeLookup[path]
		if ok {
			for _, edge := range addEdges {
				edgeId := fmt.Sprintf("%s-%s", edge.Source, edge.Target)
				_, alreadyAdded := seenEdges[edgeId]
				if !alreadyAdded {
					graph.Graph.Edges = append(graph.Graph.Edges, edge)
					seenEdges[edgeId] = true
				}
			}
		}
	}

	// addNewNode to the graph Nodes if we haven't yet
	var addNewNode = func(path string, node jgf.Node) {
		_, ok := graph.NodeMap[path]
		if !ok {
			graph.NodeMap[path] = node
			graph.Graph.Nodes = append(graph.Graph.Nodes, node)
			addNewEdges(path)
		}
	}

	// For each core path provided, add it and edges to the graph
	for _, path := range in.Cores {
		node, ok := c.nodeLookup[path]
		if !ok {
			return &response, fmt.Errorf("Cannot find node with path %s", path)
		}

		// If it isn't added to our graph node map, add it
		addNewNode(path, node)

		// Parse the entire path and add nodes up root
		parts := strings.Split(path, "/")
		for idx := 1; idx < len(parts); idx++ {
			path := strings.Join(parts[0:idx], "/")
			node := c.nodePaths[path]
			addNewNode(path, node)
		}
	}

	// Serialize the cancel request to string
	graphStr, err := graph.ToJson()
	if err != nil {
		return &response, err
	}
	fmt.Println(graphStr)

	// Now we issue a cancel request
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// This error reflects the success or failure of the cancel request
	in.Jgf = graphStr
	res, err := c.service.PartialCancel(ctx, in)

	// err is an issue with the call, res.Error is the cancel being successful or not
	if err != nil || res.Error != 0 {
		return &response, err
	}
	response.Status = pb.PartialCancelResponse_PARTIAL_CANCEL_SUCCESS
	return &response, err
}

// Cancel a job
func (c *FluxionClient) Cancel(ctx context.Context, in *pb.CancelRequest, opts ...grpc.CallOption) (*pb.CancelResponse, error) {

	response := &pb.CancelResponse{}
	if !c.Connected() {
		return response, errors.New("client is not connected")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// This error reflects the success or failure of the cancel request
	res, err := c.service.Cancel(ctx, in)
	if err != nil {
		response.Status = pb.CancelResponse_CANCEL_REQUEST_ERROR
		return response, err
	}

	// And this error is if the cancel was successful or not
	if res.Error != 0 {
		response.Status = pb.CancelResponse_CANCEL_ERROR
		return response, err
	}
	response.Status = pb.CancelResponse_CANCEL_SUCCESS
	return response, err
}

// Init the fluxion nodes
func (c *FluxionClient) Init(ctx context.Context, in *pb.InitRequest, opts ...grpc.CallOption) (*pb.InitResponse, error) {
	response := &pb.InitResponse{Status: pb.InitResponse_INIT_ERROR}

	if !c.Connected() {
		return response, errors.New("client is not connected")
	}

	// Partial cancel requires saving lookups and a view of the graph
	// that can help to create the partial cancel request.
	// Load the JGF - we keep a copy for partial cancel
	if in.EnablePartialCancel {
		graph, err := jgf.LoadFluxJGF(in.Jgf)
		if err != nil {
			return response, err
		}
		c.jgf = &graph
		c.cacheGraph()
	}

	// Contact the server...
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return c.service.Init(ctx, in)
}
