package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/converged-computing/fluxion/pkg/client"
	pb "github.com/converged-computing/fluxion/pkg/fluxion-grpc"
)

const (
	defaultPort   = "4242"
	defaultHost   = "localhost"
	defaultPolicy = "lonode"
)

//go:embed cluster-nodes-v2.json
var jgfV2 string

// Note we are using v1 that flux currently uses
//
//go:embed cluster-nodes.json
var jgf string

//go:embed jobspec.yaml
var jobspecText string

// Run will register the cluster with rainbow
func main() {
	fmt.Println("🦩️ This is the fluxion graph client")
	grpcPort := flag.String("port", defaultPort, "Port for grpc service")
	grpcHost := flag.String("host", defaultHost, "Host for grpc service")
	jsFile := flag.String("jobspec", "", "Jobspec file")
	matchPolicy := flag.String("policy", defaultPolicy, "Policy for fluxion")
	flag.Parse()

	// Ensure our port starts with :
	port := *grpcPort
	host := *grpcHost
	policy := *matchPolicy
	jobspecFile := *jsFile

	// Get the default Jobspec to read
	if jobspecFile == "" {
		log.Fatalf("Please provide a path to a --jobspec")
	}

	// Ensure the port has :
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

	// Step 1: create a cluster from the faux JGF
	request := &pb.InitRequest{Policy: policy, Jgf: jgf}
	response, err := c.Init(ctx, request)
	if err != nil {
		log.Fatalf("Issue with init request: %s", err)
	}

	// If we get here, success! Dump all the stuff.
	log.Printf("⭐️ Init cluster status: %s", response.Status)

	// Step 2: do a match!
	// Default Request here is "allocate"
	matchRequest := &pb.MatchRequest{Jobspec: jobspecText}
	matchResponse, err := c.Match(ctx, matchRequest)
	if err != nil {
		log.Fatalf("Issue with match request: %s", err)
	}
	log.Printf("⭐️ Match status: %s", matchResponse.Status)
	log.Printf("     Allocation: %s", matchResponse.Allocation)
	log.Printf("       Overhead: %f", matchResponse.Overhead)
	log.Printf("       Reserved: %t", matchResponse.Reserved)
	log.Printf("          Jobid: %d", matchResponse.Jobid)
	log.Printf("             At: %d", matchResponse.At)

	// Sleep a little to let job "run"
	fmt.Println("😴️ Sleeping for 3 seconds before cancel...")
	time.Sleep(5 * time.Second)

	cancelRequest := &pb.CancelRequest{JobID: matchResponse.Jobid}
	cancelResponse, err := c.Cancel(ctx, cancelRequest)
	if err != nil {
		log.Fatalf("Issue with cancel request: %s", err)
	}
	log.Printf("⭐️ Cancel status: %s", cancelResponse.Status)
}
