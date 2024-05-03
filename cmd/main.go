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
	"github.com/converged-computing/fluxion/pkg/service"
	svcPb "github.com/converged-computing/fluxion/pkg/service-grpc"
)

const (
	defaultPort           = "4242"
	enableExternalService = false
)

var responsechan chan string

func main() {
	fmt.Println("This is the fluxion graph server")
	grpcPort := flag.String("port", defaultPort, "Port for grpc service")
	enableServicePlugin := flag.Bool("external-service", enableExternalService, "Flag to enable the external service (defaults to false)")
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

	// External plugin (Kubectl) GRPC
	// This will eventually be an external GRPC module that can
	// be shared by fluence (flux-k8s) and fluence-kubectl
	// We give it a handle to Flux to get the state of groups
	// and job Ids. The direct interaction with Fluxion
	// happens through the other service handle
	if *enableServicePlugin {
		plugin := service.ExternalService{}
		plugin.Init()
		svcPb.RegisterExternalPluginServiceServer(s, &plugin)
	}

	fmt.Printf("[GRPCServer] gRPC Listening on %s\n", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("[GRPCServer] failed to serve: %v\n", err)
	}
	fmt.Printf("[GRPCServer] Exiting\n")
}
