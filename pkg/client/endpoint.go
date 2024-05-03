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

	// An error here is an error with making the request
	response, err := c.service.Match(ctx, in)
	if err != nil {
		fmt.Printf("[Match] did not receive any match response: %v\n", err)
		return response, err
	}
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
	response := &pb.InitResponse{}

	if !c.Connected() {
		return response, errors.New("client is not connected")
	}

	// Contact the server...
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return c.service.Init(ctx, in)
}
