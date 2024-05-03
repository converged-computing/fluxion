package fluxion

import (
	"fmt"
	"os"
	"strings"

	pb "github.com/converged-computing/fluxion/pkg/fluxion-grpc"
	"github.com/converged-computing/fluxion/pkg/jobspec"
	"github.com/flux-framework/fluxion-go/pkg/fluxcli"
	"k8s.io/klog/v2"

	"context"
	"errors"
)

type Fluxion struct {
	cli *fluxcli.ReapiClient
	pb.UnimplementedFluxcliServiceServer
}

// Init creates a new client to interaction with the fluxion API
func (f *Fluxion) Init(ctx context.Context, in *pb.InitRequest) (*pb.InitResponse, error) {
	f.cli = fluxcli.NewReapiClient()

	response := pb.InitResponse{}
	fmt.Printf("[Fluxion] Created flux memory graph")

	// Note that we currently take in JGF version 2, and convert to v1 here for flux.
	// gv1, err := createJGF(in.Jgf)
	policyRequest := "{}"
	if in.Policy != "" {
		policyRequest = string("{\"matcher_policy\": \"" + in.Policy + "\"}")
		fmt.Printf("[Fluxion] match policy: %s", policyRequest)
	}

	// Dump the jgf version 1 into a string for flux
	// raw, err := json.Marshal(gv1)
	// out := string(raw)
	// fmt.Println(out)

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

	klog.Infof("[Fluxion] received cancel request %v\n", in)
	err := s.cli.Cancel(int64(in.JobID), true)
	if err != nil {
		return nil, errors.New("Error in Cancel")
	}

	// Why would we have an error code here if we check above?
	// This (I think) should be an error code for the specific job
	dr := &pb.CancelResponse{JobID: in.JobID}
	klog.Infof("[Fluxion] sending cancel response %v\n", dr)
	klog.Infof("[Fluxion] cancel errors so far: %s\n", s.cli.GetErrMsg())

	reserved, at, overhead, mode, fluxerr := s.cli.Info(int64(in.JobID))
	klog.Infof("\n\t----Job Info output---")
	klog.Infof("jobid: %d\nreserved: %t\nat: %d\noverhead: %f\nmode: %s\nerror: %d\n", in.JobID, reserved, at, overhead, mode, fluxerr)

	klog.Infof("[GRPCServer] Sending Cancel response %v\n", dr)
	return dr, nil
}

// generateJobSpec generates a jobspec for a match request and returns the string
// TODO this should be updated to use jobspec-go, and not need a special
// library here
func (s *Fluxion) generateJobspec(in *pb.MatchRequest) ([]byte, error) {

	spec := []byte{}

	// Create a temporary file to write and read the jobspec
	// The first parameter here as the empty string creates in /tmp
	file, err := os.CreateTemp("", "jobspec.*.yaml")
	if err != nil {
		return spec, err
	}
	defer os.Remove(file.Name())
	jobspec.CreateJobSpecYaml(in.Resources, in.Count, file.Name())

	spec, err = os.ReadFile(file.Name())
	if err != nil {
		return spec, errors.New("Error reading jobspec")
	}
	return spec, err
}

// Match wraps the MatchAllocate function of the fluxion go bindings
// If a match is not possible, we return the error and an empty response
func (s *Fluxion) Match(ctx context.Context, in *pb.MatchRequest) (*pb.MatchResponse, error) {

	emptyResponse := &pb.MatchResponse{}

	// Prepare an empty match response (that can still be serialized)
	klog.Infof("[Fluxion] Received Match request %v\n", in)

	// Generate the jobspec, written to temporary file and read as string
	spec, err := s.generateJobspec(in)
	if err != nil {
		return emptyResponse, err
	}

	// Ask flux to match allocate!
	reserved, allocated, at, overhead, jobid, fluxerr := s.cli.MatchAllocate(false, string(spec))

	// TODO get rid of this or refactor
	printOutput(reserved, allocated, at, overhead, jobid, fluxerr)

	// Be explicit about errors (or not)
	errorMessages := s.cli.GetErrMsg()
	if errorMessages == "" {
		klog.Infof("[Fluxion] There are no errors")
	} else {
		klog.Infof("[Fluxion] Match errors so far: %s\n", errorMessages)
	}
	if fluxerr != nil {
		klog.Infof("[Fluxion] Match Flux err is %w\n", fluxerr)
		return emptyResponse, errors.New("[Fluxion] Error in ReapiCliMatchAllocate")
	}

	// This usually means we cannot allocate
	// We need to return an error here otherwise we try to pass an empty string
	// to other RPC endpoints and get back an error.
	if allocated == "" {
		klog.Infof("[Fluxion] Allocated is empty")
		return emptyResponse, errors.New("Allocation was not possible")
	}

	// Pass the spec name in so we can include it in the allocation result
	// This will allow us to inspect the ordering later.
	nodetasks := parseAllocResult(allocated)
	nodetaskslist := make([]*pb.NodeAlloc, len(nodetasks))
	for i, result := range nodetasks {
		nodetaskslist[i] = &pb.NodeAlloc{
			NodeID: result.Basename,
			Tasks:  int32(result.CoreCount) / in.Resources.Cpu,
		}
	}
	mr := &pb.MatchResponse{PodID: in.Resources.Id, Nodelist: nodetaskslist, JobID: int64(jobid)}
	klog.Infof("[Fluxion] Match response %v \n", mr)
	return mr, nil
}
