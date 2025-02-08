package fluxion

import (
	"fmt"
	"strings"

	pb "github.com/converged-computing/fluxion/pkg/fluxion-grpc"
	"github.com/flux-framework/fluxion-go/pkg/fluxcli"
	"github.com/flux-framework/fluxion-go/pkg/types"

	"context"
	"errors"
)

type Fluxion struct {
	cli *fluxcli.ReapiClient
	pb.UnimplementedFluxionServiceServer
}

// Init creates a new client to interaction with the fluxion API
func (f *Fluxion) Init(ctx context.Context, in *pb.InitRequest) (*pb.InitResponse, error) {
	f.cli = fluxcli.NewReapiClient()

	response := pb.InitResponse{}
	fmt.Printf("[Fluxion] Created flux memory graph")

	// Update policy if set
	policyRequest := "{}"
	if in.Policy != "" {
		policyRequest = string("{\"matcher_policy\": \"" + in.Policy + "\"}")
		fmt.Printf("[Fluxion] match policy: %s", policyRequest)
	}

	err := f.cli.InitContext(in.Jgf, policyRequest)
	if err != nil {
		response.Status = pb.InitResponse_INIT_ERROR
		f.ShowError()
		return &response, err
	}

	// Successful response!
	response.Status = pb.InitResponse_INIT_SUCCESS
	if err != nil {
		response.Status = pb.InitResponse_INIT_ERROR
		f.ShowError()
	}
	return &response, err
}

// Get error wrapped in more meaningful message
func (f *Fluxion) ShowError() {
	msg := f.cli.GetErrMsg()
	fmt.Println("\n" + strings.ReplaceAll(msg, "\n\n", ""))
}

// Cancel wraps the Cancel function of the fluxion go bindings
func (s *Fluxion) Cancel(ctx context.Context, in *pb.CancelRequest) (*pb.CancelResponse, error) {

	response := pb.CancelResponse{}
	fmt.Printf("[Fluxion] received cancel request %v\n", in)
	err := s.cli.Cancel(in.JobID, true)
	if err != nil {
		fmt.Printf("[Fluxion] issue with cancel %s\n", err)
		response.Status = pb.CancelResponse_CANCEL_ERROR
		return &response, err
	}
	response.Status = pb.CancelResponse_CANCEL_SUCCESS
	return &response, err
}

// PartialCancel allows for adding additional JGF to "prune" or partially cancel
func (s *Fluxion) PartialCancel(ctx context.Context, in *pb.PartialCancelRequest) (*pb.PartialCancelResponse, error) {

	response := pb.PartialCancelResponse{}
	fmt.Printf("[Fluxion] received partial cancel request %v\n", in)

	// The last argument is noent_ok, ok if doesn't exist
	fullRemoval, err := s.cli.PartialCancel(in.JobID, in.Jgf, in.MustExist)

	if err != nil {
		fmt.Printf("[Fluxion] issue with partial cancel %s\n", err)
		response.Status = pb.PartialCancelResponse_PARTIAL_CANCEL_ERROR
		return &response, err
	}
	response.Status = pb.PartialCancelResponse_PARTIAL_CANCEL_SUCCESS
	response.FullRemoval = fullRemoval
	return &response, nil
}

// Match wraps the MatchAllocate function of the fluxion go bindings
// If a match is not possible, we return the error and an empty response
func (s *Fluxion) Match(ctx context.Context, in *pb.MatchRequest) (*pb.MatchResponse, error) {

	response := &pb.MatchResponse{Status: pb.MatchResponse_MATCH_ERROR}

	// Ask flux to match allocate!
	reserved, allocated, at, overhead, jobid, fluxerr := s.cli.MatchAllocate(in.Reservation, in.Jobspec)

	// Be explicit about errors (or not)
	errorMessages := s.cli.GetErrMsg()
	if errorMessages != "" {
		fmt.Println("[Fluxion] Match errors so far: %s\n", errorMessages)
	}
	if fluxerr != nil {
		fmt.Printf("[Fluxion] Match Flux err is %s\n", fluxerr)
		return response, errors.New("[Fluxion] Error in ReapiCliMatchAllocate")
	}

	// This usually means we cannot allocate
	// We need to return an error here otherwise we try to pass an empty string
	// to other RPC endpoints and get back an error.
	if allocated == "" {
		fmt.Printf("[Fluxion] Allocated is empty")
		return response, errors.New("Allocation was not possible")
	}

	// Return the raw match response - this could be better parsed
	response.Status = pb.MatchResponse_MATCH_SUCCESS
	response.Allocation = allocated
	response.Jobid = int64(jobid)
	response.Reserved = reserved
	response.At = at
	response.Overhead = float32(overhead)
	printMatch(response)
	return response, nil
}

// Satisfy wraps the MatchSatisfiability fluxion endpoint
func (s *Fluxion) Satisfy(ctx context.Context, in *pb.SatisfyRequest) (*pb.SatisfyResponse, error) {

	response := &pb.SatisfyResponse{Status: pb.SatisfyResponse_SATISFY_ERROR}
	reserved, allocated, at, overhead, jobid, fluxerr := s.cli.Match(types.MatchSatisfiability, in.Jobspec)

	// Be explicit about errors (or not)
	errorMessages := s.cli.GetErrMsg()
	if errorMessages != "" {
		fmt.Println("[Fluxion] Satisfy errors so far: %s\n", errorMessages)
	}
	if fluxerr != nil {
		fmt.Printf("[Fluxion] Satisfy Flux err is %s\n", fluxerr)
		return response, errors.New("[Fluxion] Error in MatchSatisfy")
	}
	if allocated == "" {
		fmt.Printf("[Fluxion] Allocated is empty")
		return response, errors.New("Allocation for satisfy was not possible")
	}

	// Return the raw match response - this could be better parsed
	response.Status = pb.SatisfyResponse_SATISFY_SUCCESS
	response.Allocation = allocated
	response.Jobid = int64(jobid)
	response.Reserved = reserved
	response.At = at
	response.Overhead = float32(overhead)
	printSatisfy(response)
	return response, nil
}

func printMatch(response *pb.MatchResponse) {
	fmt.Println("\n💼️ Allocation Match Result")
	fmt.Printf("       Overhead: %f\n", response.Overhead)
	fmt.Printf("       Reserved: %t\n", response.Reserved)
	fmt.Printf("         Status: %s\n", response.Status)
	fmt.Printf("          Jobid: %d\n", response.Jobid)
	fmt.Printf("             At: %d\n", response.At)
	fmt.Printf("     Allocation: %s\n", response.Allocation)
}

func printSatisfy(response *pb.SatisfyResponse) {
	fmt.Println("\n💼️ Allocation Satisfy Result")
	fmt.Printf("       Overhead: %f\n", response.Overhead)
	fmt.Printf("       Reserved: %t\n", response.Reserved)
	fmt.Printf("         Status: %s\n", response.Status)
	fmt.Printf("          Jobid: %d\n", response.Jobid)
	fmt.Printf("             At: %d\n", response.At)
	fmt.Printf("     Allocation: %s\n", response.Allocation)
}
