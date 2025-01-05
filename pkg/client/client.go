package client

import (
	"context"
	"log"

	pb "github.com/converged-computing/fluxion/pkg/fluxion-grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// FluxionClient interacts with Fluxion
type FluxionClient struct {
	host       string
	connection *grpc.ClientConn
	service    pb.FluxionServiceClient
}

var _ Client = (*FluxionClient)(nil)

// Client interface defines functions required for a valid client
type Client interface {
	Match(ctx context.Context, in *pb.MatchRequest, opts ...grpc.CallOption) (*pb.MatchResponse, error)
	Cancel(ctx context.Context, in *pb.CancelRequest, opts ...grpc.CallOption) (*pb.CancelResponse, error)
	Init(ctx context.Context, in *pb.InitRequest, opts ...grpc.CallOption) (*pb.InitResponse, error)

	// Functions that aren't related to fluxion directly
	Close() error
	GetHost() string
	Connected() bool
}

// NewClient creates a new FluxionClient
func NewClient(host string) (Client, error) {
	if host == "" {
		return nil, errors.New("host is required")
	}

	log.Printf("🦩️ starting client (%s)...", host)
	c := &FluxionClient{host: host}

	// Set up a connection to the server.
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(c.GetHost(), creds, grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to connect to %s", host)
	}

	c.connection = conn
	c.service = pb.NewFluxionServiceClient(conn)
	return c, nil
}

// Close closes the created resources (e.g. connection).
func (c *FluxionClient) Close() error {
	if c.connection != nil {
		return c.connection.Close()
	}
	return nil
}

// Connected returns  true if we are connected and the connection is ready
func (c *FluxionClient) Connected() bool {
	return c.service != nil && c.connection != nil && c.connection.GetState() == connectivity.Ready
}

// GetHost returns the private hostn name
func (c *FluxionClient) GetHost() string {
	return c.host
}
