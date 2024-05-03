package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/converged-computing/fluxion/pkg/fluxion"
	pb "github.com/converged-computing/fluxion/pkg/fluxion-grpc"
)

const (
	defaultPort = "4242"
)

var responsechan chan string

func main() {
	fmt.Println("This is the fluxion graph server")
	grpcPort := flag.String("port", defaultPort, "Port for grpc service")
	flag.Parse()

	// Ensure our port starts with :
	port := *grpcPort
	if !strings.HasPrefix(":", port) {
		port = fmt.Sprintf(":%s", port)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("[GRPCServer] failed to listen: %v\n", err)
	}

	// Fluxion GRPC
	// We don't init here, that is done via gRPC
	flux := fluxion.Fluxion{}
	responsechan = make(chan string)
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)
	pb.RegisterFluxcliServiceServer(s, &flux)
	fmt.Printf("[GRPCServer] gRPC Listening on %s\n", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("[GRPCServer] failed to serve: %v\n", err)
	}
	fmt.Printf("[GRPCServer] Exiting\n")
}
