package client

import (
	"context"
	"fmt"
	"time"

	pb "github.com/converged-computing/fluxion/pkg/fluxion-grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func (c *FluxionClient) Match(ctx context.Context, in *pb.MatchRequest, opts ...grpc.CallOption) (*pb.MatchResponse, error) {
	response := &pb.MatchResponse{}
	if !c.Connected() {
		return response, errors.New("client is not connected")
	}

	// If no request type provided, assume allocate
	if in.Request == "" {
		in.Request = "allocate"
	}
	// Now contact the rainbow server with clusters...
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	request := &pb.MatchRequest{
		Resources: in.Resources,
		Request:   in.Request,
		Count:     in.Resources.Count,
	}

	// An error here is an error with making the request
	r, err := c.service.Match(context.Background(), request)
	if err != nil {
		fmt.Printf("[Match] did not receive any match response: %v\n", err)
		return response, err
	}

	fmt.Printf("[Match] Match response ID %s\n", r.GetPodID())

	// Get the nodelist and inspect
	nodelist := r.GetNodelist()

	// Show list of nodes to user
	nodes := []string{}
	for _, node := range nodelist {
		nodes = append(nodes, node.NodeID)
	}
	jobid := uint64(r.GetJobID())
	fmt.Printf("[Match] parsed node pods list %s for job id %d\n", nodes, jobid)

	response.Nodelist = nodelist
	response.JobID = int64(jobid)
	return response, err
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
	request := &pb.CancelRequest{JobID: in.JobID}
	res, err := c.service.Cancel(context.Background(), request)
	if err != nil {
		response.Status = pb.CancelResponse_CANCEL_REQUEST_ERROR
		return response, err
	}

	// And this error is if the cancel was successful or not
	if res.Error == 0 {
		response.Status = pb.CancelResponse_CANCEL_ERROR
		return response, err
	}
	response.Status = pb.CancelResponse_CANCEL_SUCCESS
	return response, err
}

// Init the fluxion nodes
func (c *FluxionClient) Init(ctx context.Context, in *pb.InitRequest, opts ...grpc.CallOption) (*pb.InitResponse, error) {
	response := &pb.InitResponse{}

	if !c.Connected() {
		return response, errors.New("client is not connected")
	}

	// Contact the server...
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return c.service.Init(ctx, in)
}
