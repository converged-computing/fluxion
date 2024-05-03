package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/converged-computing/fluxion/pkg/client"
	pb "github.com/converged-computing/fluxion/pkg/fluxion-grpc"
)

const (
	defaultPort   = "4242"
	defaultHost   = "localhost"
	defaultPolicy = "lonode"
)

//go:embed cluster-nodes.json
var jgf string

// Run will register the cluster with rainbow
func main() {
	fmt.Println("🦩️ This is the fluxion graph client")
	grpcPort := flag.String("port", defaultPort, "Port for grpc service")
	grpcHost := flag.String("host", defaultHost, "Host for grpc service")
	matchPolicy := flag.String("policy", defaultHost, "Policy for fluxion")
	flag.Parse()

	// Ensure our port starts with :
	port := *grpcPort
	host := *grpcHost
	policy := *matchPolicy

	if !strings.HasPrefix(":", port) {
		port = fmt.Sprintf(":%s", port)
	}
	host = fmt.Sprintf("%s%s", host, port)

	// This is the host where fluxion is running, will be localhost for local development
	c, err := client.NewClient(host)
	if err != nil {
		log.Fatalf("Issue with creating client: %s", err)
	}

	// Context for subsequent requests
	ctx := context.TODO()

	// 1. Step 1: create a cluster from the faux JGF
	request := &pb.InitRequest{Policy: policy, Jgf: jgf}
	response, err := c.Init(ctx, request)
	if err != nil {
		log.Fatalf("Issue with init request: %s", err)
	}

	// If we get here, success! Dump all the stuff.
	log.Printf("status: %s", response.Status)
}
