package utils

import (
	"fmt"

	"encoding/json"

	"github.com/converged-computing/fluxion/pkg/jgf"
)

// CreateJGF creates the Json Graph Format for Fluxion
func CreateJGF(graphstring, clusterName string) (string, error) {

	// Create a Flux Json Graph Format (JGF) with all cluster nodes
	fluxgraph := jgf.InitJGF()
	cluster := fluxgraph.MakeCluster(clusterName)
	fmt.Println(cluster)

	// Rack needs to be disabled when using subnets
	// rack := fluxgraph.MakeRack(0)

	// fluxgraph.MakeEdge(cluster, rack, "contains")
	// fluxgraph.MakeEdge(rack, cluster, "in")
	// TODO read in graph here
	fmt.Println(graphstring)
	/*
		vcores := 0
		fmt.Println("Number nodes ", len(nodes.Items))
		var totalAllocCpu int64
		totalAllocCpu = 0
		sdnCount := 0

		for nodeIndex, node := range nodes.Items {

			// We should not be scheduling to the control plane
			_, ok := node.Labels[controlPlaneLabel]
			if ok {
				fmt.Println("Skipping control plane node ", node.GetName())
				continue
			}

			// Anything labeled with "skipLabel" meaning it is present,
			// should be skipped
			if *skipLabel != "" {
				_, ok := node.Labels[*skipLabel]
				if ok {
					fmt.Printf("Skipping node %s\n", node.GetName())
					continue
				}
			}

			if node.Spec.Unschedulable {
				fmt.Printf("Skipping node %s, unschedulable\n", node.GetName())
				continue
			}

			// Check if subnet already exists
			// Here we build subnets according to topology.kubernetes.io/zone label
			subnetName := node.Labels["topology.kubernetes.io/zone"]
			subnet := fluxgraph.MakeSubnet(sdnCount, subnetName)
			sdnCount = sdnCount + 1
			fluxgraph.MakeEdge(cluster, subnet, "contains")
			fluxgraph.MakeEdge(subnet, cluster, "in")

			// These are requests for existing pods, for cpu and memory
			reqs := computeTotalRequests(pods)
			cpuReqs := reqs[corev1.ResourceCPU]
			memReqs := reqs[corev1.ResourceMemory]

			// Actual values that we have available (minus requests)
			totalCpu := node.Status.Allocatable.Cpu().MilliValue()
			totalMem := node.Status.Allocatable.Memory().Value()

			// Values accounting for requests
			availCpu := int64((totalCpu - cpuReqs.MilliValue()) / 1000)
			availMem := totalMem - memReqs.Value()

			// Show existing to compare to
			fmt.Printf("\n📦️ %s\n", node.GetName())
			fmt.Printf("      allocated cpu: %d\n", cpuReqs.Value())
			fmt.Printf("      allocated mem: %d\n", memReqs.Value())
			fmt.Printf("      available cpu: %d\n", availCpu)
			fmt.Printf("       running pods: %d\n", len(pods.Items))

			// keep track of overall total
			totalAllocCpu += availCpu
			fmt.Printf("      available mem: %d\n", availMem)
			gpuAllocatable, hasGpuAllocatable := node.Status.Allocatable["nvidia.com/gpu"]

			// reslist := node.Status.Allocatable
			// resources := make([]corev1.ResourceName, 0, len(reslist))
			// for resource := range reslist {
			// 	fmt.Println("resource ", resource)
			// 	resources = append(resources, resource)
			// }
			// for _, resource := range resources {
			// 	value := reslist[resource]

			// 	fmt.Printf(" %s:\t%s\n", resource, value.String())
			// }

			workernode := fluxgraph.MakeNode(nodeIndex, false, node.Name)
			fluxgraph.MakeEdge(subnet, workernode, "contains") // this is rack otherwise
			fluxgraph.MakeEdge(workernode, subnet, "in")       // this is rack otherwise

			// socket := fluxgraph.MakeSocket(0, "socket")
			// fluxgraph.MakeEdge(workernode, socket, "contains")
			// fluxgraph.MakeEdge(socket, workernode, "in")

			if hasGpuAllocatable {
				fmt.Println("GPU Resource quantity ", gpuAllocatable.Value())
				//MakeGPU(index int, name string, size int) string {
				for index := 0; index < int(gpuAllocatable.Value()); index++ {
					gpu := fluxgraph.MakeGPU(index, "nvidiagpu", 1)
					fluxgraph.MakeEdge(workernode, gpu, "contains") // workernode was socket
					fluxgraph.MakeEdge(gpu, workernode, "in")
				}

			}

			for index := 0; index < int(availCpu); index++ {
				// MakeCore(index int, name string)
				core := fluxgraph.MakeCore(index, "core")
				fluxgraph.MakeEdge(workernode, core, "contains") // workernode was socket
				fluxgraph.MakeEdge(core, workernode, "in")

				// Question from Vanessa:
				// How can we get here and have vcores ever not equal to zero?
				if vcores == 0 {
					fluxgraph.MakeNFDProperties(core, index, "cpu-", &node.Labels)
					// fluxgraph.MakeNFDProperties(core, index, "netmark-", &node.Labels)
				} else {
					for vc := 0; vc < vcores; vc++ {
						vcore := fluxgraph.MakeVCore(core, vc, "vcore")
						fluxgraph.MakeNFDProperties(vcore, index, "cpu-", &node.Labels)
					}
				}
			}

			// MakeMemory(index int, name string, unit string, size int)
			fractionMem := availMem >> 30
			// fractionmem := (totalmem/totalcpu) >> 20
			// fmt.Println("Creating ", fractionmem, " vertices with ", 1<<10, " MB of mem")
			for i := 0; i < /*int(totalcpu) int(fractionMem); i++ {
				mem := fluxgraph.MakeMemory(i, "memory", "MB", int(1<<10))
				fluxgraph.MakeEdge(workernode, mem, "contains")
				fluxgraph.MakeEdge(mem, workernode, "in")
			}
		}*/
	totalAllocCpu := 0
	fmt.Printf("\nCan request at most %d exclusive cpu", totalAllocCpu)
	return "", nil
}

type allocation struct {
	Type      string
	Name      string
	Basename  string
	CoreCount int
}

// ParseAllocResult takes an allocated (string) and parses into a list of allocation
// We include the pod namespace/name for debugging later
func ParseAllocResult(allocated string) []allocation {
	var dat map[string]interface{}
	result := []allocation{}

	// Keep track of total core count across allocated
	corecount := 0

	// This should not happen - the string we get back should parse.
	if err := json.Unmarshal([]byte(allocated), &dat); err != nil {
		panic(err)
	}
	// Parse graph and nodes into interfaces
	// TODO look at github.com/mitchellh/mapstructure
	// that might make this easier
	nodes := dat["graph"]
	str1 := nodes.(map[string]interface{})
	str2 := str1["nodes"].([]interface{})

	for _, item := range str2 {
		str1 = item.(map[string]interface{})
		metadata := str1["metadata"].(map[string]interface{})
		if metadata["type"].(string) == "core" {
			corecount = corecount + 1
		}
		if metadata["type"].(string) == "node" {
			result = append(result, allocation{
				Type:      metadata["type"].(string),
				Name:      metadata["name"].(string),
				Basename:  metadata["basename"].(string),
				CoreCount: corecount,
			})

			// Reset the corecount once we've added to a node
			corecount = 0
		}
	}
	fmt.Println("Final node result")
	for i, alloc := range result {
		fmt.Printf("Node %d: %s\n", i, alloc.Name)
		fmt.Printf("  Type: %s\n  Name: %s\n  Basename: %s\n  CoreCount: %d\n",
			alloc.Type, alloc.Name, alloc.Basename, alloc.CoreCount)

	}
	return result
}

// Utility functions
func PrintOutput(reserved bool, allocated string, at int64, overhead float64, jobid uint64, fluxerr error) {
	fmt.Println("\n\t----Match Allocate output---")
	fmt.Printf("jobid: %d\nreserved: %t\nallocated: %s\nat: %d\noverhead: %f\n", jobid, reserved, allocated, at, overhead)

	// Only print error if we had one
	if fluxerr != nil {
		fmt.Printf("error: %w\n", fluxerr)
	}
}
