package client

import (
	"context"
	"log"

	"github.com/compspec/jgf-go/pkg/jgf"
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

	// These are intended for partial cancel
	jgf *jgf.FluxJGF

	nodeLookup map[string]jgf.Node

	// Store nodes based on paths
	nodePaths  map[string]jgf.Node
	edgeLookup map[string][]jgf.Edge
	hostLookup map[string]string
}

var _ Client = (*FluxionClient)(nil)

// Client interface defines functions required for a valid client
type Client interface {
	Match(ctx context.Context, in *pb.MatchRequest, opts ...grpc.CallOption) (*pb.MatchResponse, error)
	Satisfy(ctx context.Context, in *pb.SatisfyRequest, opts ...grpc.CallOption) (*pb.SatisfyResponse, error)
	Cancel(ctx context.Context, in *pb.CancelRequest, opts ...grpc.CallOption) (*pb.CancelResponse, error)
	PartialCancel(ctx context.Context, in *pb.PartialCancelRequest, opts ...grpc.CallOption) (*pb.PartialCancelResponse, error)
	Init(ctx context.Context, in *pb.InitRequest, opts ...grpc.CallOption) (*pb.InitResponse, error)

	// Functions that aren't related to fluxion directly
	Close() error
	GetHost() string
	Connected() bool
	cacheGraph()
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
	conn, err := grpc.Dial(c.GetHost(), creds)
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
