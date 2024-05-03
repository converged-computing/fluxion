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
	defaultHost = ""
)

var responsechan chan string

func main() {
	fmt.Println("🦩️ This is the fluxion graph server")
	grpcPort := flag.String("port", defaultPort, "Port for grpc service")
	grpcHost := flag.String("host", defaultHost, "Host for grpc service")
	flag.Parse()

	// Ensure our port starts with :
	port := *grpcPort
	host := *grpcHost
	if !strings.HasPrefix(":", port) {
		port = fmt.Sprintf(":%s", port)
	}

	address := fmt.Sprintf("%s%s", host, port)
	lis, err := net.Listen("tcp", address)
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
	pb.RegisterFluxionServiceServer(s, &flux)
	fmt.Printf("[GRPCServer] gRPC Listening on %s\n", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("[GRPCServer] failed to serve: %v\n", err)
	}
	fmt.Printf("[GRPCServer] Exiting\n")
}
