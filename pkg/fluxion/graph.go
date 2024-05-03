package fluxion

import (
	"encoding/json"
	"fmt"

	v1 "github.com/converged-computing/jsongraph-go/jsongraph/v1/graph"
	"github.com/converged-computing/jsongraph-go/jsongraph/v2/graph"
)

// ReadNodeJsonGraph reads in the node JGF
func readNodeJsonGraph(raw string) (graph.JsonGraph, error) {
	g := graph.JsonGraph{}
	err := json.Unmarshal([]byte(raw), &g)
	if err != nil {
		return g, fmt.Errorf("error unmarshalling %s", err)
	}
	return g, nil
}

// createJGF parses the raw string into a proper JGF struct
// We start with v2 and parse into v1 for flux
func createJGF(raw string) (*v1.JsonGraph, error) {

	// We will return JGF version 1
	gv1 := v1.NewGraph()

	// This reads in v2
	g, err := readNodeJsonGraph(raw)
	if err != nil {
		return gv1, err
	}

	// And we parse into v1
	for _, edge := range g.Graph.Edges {
		newEdge := convertEdge(edge)
		gv1.Graph.Edges = append(gv1.Graph.Edges, newEdge)
	}
	for nid, node := range g.Graph.Nodes {
		newNode := convertNode(nid, node)
		gv1.Graph.Nodes = append(gv1.Graph.Nodes, newNode)
	}
	return gv1, nil
}

// convertEdge converts a v2 edge to a v1 edge (backwards)
// This would be easier if these had the same underlying type, but I prefer
// the separation in the typing for different versions
func convertEdge(edge graph.Edge) v1.Edge {
	return v1.Edge{
		Id:       edge.Id,
		Source:   edge.Source,
		Target:   edge.Target,
		Relation: edge.Relation,
		Label:    edge.Label,
		Directed: edge.Directed,
		Metadata: edge.Metadata,
	}
}

// convertNode converts a v2 node to a v1 node.
func convertNode(nid string, node graph.Node) v1.Node {
	return v1.Node{
		Id:       nid,
		Label:    node.Label,
		Metadata: node.Metadata,
	}
}
