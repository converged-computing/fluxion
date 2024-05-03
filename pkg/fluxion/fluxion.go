package fluxion

import (
	"fmt"
	"os"

	pb "github.com/converged-computing/fluxion/pkg/fluxion-grpc"
	"github.com/converged-computing/fluxion/pkg/jobspec"
	"github.com/converged-computing/fluxion/pkg/utils"
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
	fmt.Printf("[Fluence] Created flux resource client %s", f.cli)

	// Note that we currently take in JGF version 2, and convert to v1 here for flux.
	// This is intentional so when flux switches to v2 it should be easy to do here.
	jgf, err := utils.CreateJGF(in.Jgf, in.ClusterName)
	if err != nil {
		response.Status = pb.InitResponse_INIT_ERROR
		return &response, err
	}

	p := "{}"
	if in.Policy != "" {
		p = string("{\"matcher_policy\": \"" + in.Policy + "\"}")
		fmt.Printf("[Fluence] match policy: %s", p)
	}

	f.cli.InitContext(string(jgf), p)

	// Successful response!
	response.Status = pb.InitResponse_INIT_SUCCESS
	return &response, err
}

// Cancel wraps the Cancel function of the fluxion go bindings
func (s *Fluxion) Cancel(ctx context.Context, in *pb.CancelRequest) (*pb.CancelResponse, error) {

	klog.Infof("[Fluence] received cancel request %v\n", in)
	err := s.cli.Cancel(int64(in.JobID), true)
	if err != nil {
		return nil, errors.New("Error in Cancel")
	}

	// Why would we have an error code here if we check above?
	// This (I think) should be an error code for the specific job
	dr := &pb.CancelResponse{JobID: in.JobID}
	klog.Infof("[Fluence] sending cancel response %v\n", dr)
	klog.Infof("[Fluence] cancel errors so far: %s\n", s.cli.GetErrMsg())

	reserved, at, overhead, mode, fluxerr := s.cli.Info(int64(in.JobID))
	klog.Infof("\n\t----Job Info output---")
	klog.Infof("jobid: %d\nreserved: %t\nat: %d\noverhead: %f\nmode: %s\nerror: %d\n", in.JobID, reserved, at, overhead, mode, fluxerr)

	klog.Infof("[GRPCServer] Sending Cancel response %v\n", dr)
	return dr, nil
}

// generateJobSpec generates a jobspec for a match request and returns the string
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
	klog.Infof("[Fluence] Received Match request %v\n", in)

	// Generate the jobspec, written to temporary file and read as string
	spec, err := s.generateJobspec(in)
	if err != nil {
		return emptyResponse, err
	}

	// Ask flux to match allocate!
	reserved, allocated, at, overhead, jobid, fluxerr := s.cli.MatchAllocate(false, string(spec))
	utils.PrintOutput(reserved, allocated, at, overhead, jobid, fluxerr)

	// Be explicit about errors (or not)
	errorMessages := s.cli.GetErrMsg()
	if errorMessages == "" {
		klog.Infof("[Fluence] There are no errors")
	} else {
		klog.Infof("[Fluence] Match errors so far: %s\n", errorMessages)
	}
	if fluxerr != nil {
		klog.Infof("[Fluence] Match Flux err is %w\n", fluxerr)
		return emptyResponse, errors.New("[Fluence] Error in ReapiCliMatchAllocate")
	}

	// This usually means we cannot allocate
	// We need to return an error here otherwise we try to pass an empty string
	// to other RPC endpoints and get back an error.
	if allocated == "" {
		klog.Infof("[Fluence] Allocated is empty")
		return emptyResponse, errors.New("Allocation was not possible")
	}

	// Pass the spec name in so we can include it in the allocation result
	// This will allow us to inspect the ordering later.
	nodetasks := utils.ParseAllocResult(allocated)
	nodetaskslist := make([]*pb.NodeAlloc, len(nodetasks))
	for i, result := range nodetasks {
		nodetaskslist[i] = &pb.NodeAlloc{
			NodeID: result.Basename,
			Tasks:  int32(result.CoreCount) / in.Resources.Cpu,
		}
	}
	mr := &pb.MatchResponse{PodID: in.Resources.Id, Nodelist: nodetaskslist, JobID: int64(jobid)}
	klog.Infof("[Fluence] Match response %v \n", mr)
	return mr, nil
}
