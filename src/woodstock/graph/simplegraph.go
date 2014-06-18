package graph


// SimpleGraphNode is the node structure that is stored in SimpleGraph. It also
// stores the different edges which are connected to it.
type SimpleGraphNode struct {
	id    int
	edges map[int]*SimpleGraphEdge
}

func NewSimpleGraphNode() *SimpleGraphNode {
	return &SimpleGraphNode{edges: map[int]*SimpleGraphEdge{}}
}

// SimpleGraphEdge is the edge structure thats is stored in SimpleGraph and that
// is stores Nodes connected to it in a list
type SimpleGraphEdge struct {
	id int
	// First one is the input Node and second one is the output Node
	nodes []*SimpleGraphNode
}

func NewSimpleGraphEdge() *SimpleGraphEdge {
	return &SimpleGraphEdge{nodes: []*SimpleGraphNode{}}
}

// SimpleGraph is the implementation of a simple adjacency list graph that
// where links are stored in a list and that you can get nodes from edges
type SimpleGraph struct {
	nodes map[int]*SimpleGraphNode
	edges map[int]*SimpleGraphEdge

	lastEdgeID int
	lastNodeID int
}

func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{
		lastEdgeID: -1,
		lastNodeID: -1,
		nodes:      map[int]*SimpleGraphNode{},
		edges:      map[int]*SimpleGraphEdge{}}
}

// AddNode adds a node to the SimpleGraph and returns it
func (g SimpleGraph) AddNode() *SimpleGraphNode {
	n := NewSimpleGraphNode()

	id := g.lastNodeID + 1
	n.id = id
	g.lastNodeID = id

	g.nodes[id] = n
	return n
}

// AddEdge adds an edge to the SimpleGraph connected to fromNode and toNode
func (g SimpleGraph) AddEdge(fromNode *SimpleGraphNode,
	toNode *SimpleGraphNode) *SimpleGraphEdge {

	id := g.lastEdgeID + 1
	g.lastEdgeID = id

	e := NewSimpleGraphEdge()
	e.id = id
	e.nodes = append(e.nodes, fromNode)
	e.nodes = append(e.nodes, toNode)

	fromNode.edges[id] = e
	toNode.edges[id] = e

	g.edges[id] = e
	return e
}
// RemoveNode removes n from SimpleGraph, if there are edges connected to that
// node they are also destroyed
func (g SimpleGraph) RemoveNode(n *SimpleGraphNode) (err error) {

	delete(g.nodes, n.id)

	for _, e := range n.edges {
		g.RemoveEdge(e)
	}

	return nil
}
// RemoveEdge removes the edge e from the SimpleGraph and deletes the link to it
// from the nodes it connects together
func (g SimpleGraph) RemoveEdge(e *SimpleGraphEdge) (err error) {

	delete(g.edges, e.id)

	for _, n := range e.nodes {
		delete(n.edges, e.id)
	}

	return nil
}
