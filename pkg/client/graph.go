package client

import (
	"github.com/compspec/jgf-go/pkg/jgf"
)

// The edge lookup will allow us to add connected nodes
// We need to be able to map a node path to a list of edges
// The node path gets us the node id (source)
func (c *FluxionClient) addEdge(node *jgf.Node, edge *jgf.Edge) {
	path := node.Metadata.Paths["containment"]
	_, ok := c.edgeLookup[path]
	if !ok {
		c.edgeLookup[path] = []jgf.Edge{}
	}
	c.edgeLookup[path] = append(c.edgeLookup[path], *edge)
}

// cacheGraph saves lookups for edges and nodes for later (more efficient)
// generation of partial cancel requests.
func (c *FluxionClient) cacheGraph() {

	graph := c.jgf.Graph

	// Parse nodes first so we can match the containment path to the host
	for _, node := range graph.Nodes {
		c.nodeLookup[node.Id] = node
		c.nodePaths[node.Metadata.Paths["containment"]] = node
	}

	for _, edge := range graph.Edges {
		targetNode := c.nodeLookup[edge.Target]
		sourceNode := c.nodeLookup[edge.Source]
		c.addEdge(&targetNode, &edge)
		c.addEdge(&sourceNode, &edge)
	}

	// We need to be able to easily look up the host name of
	// a node based on the graph path
	for _, node := range graph.Nodes {
		nodePath := node.Metadata.Paths["containment"]
		c.nodeLookup[node.Id] = node
		if node.Metadata.Type == "node" {
			c.hostLookup[nodePath] = node.Metadata.Basename
		}
	}
}
