package graph

// SimpleGraphNode is the node structure that is stored in SimpleGraph. It also
// stores the different edges which are connected to it.
type SimpleGraphNode struct {
	id int

	Edges map[int]*SimpleGraphEdge

	Attributes map[string]string
}

func NewSimpleGraphNode() *SimpleGraphNode {
	return &SimpleGraphNode{Edges: map[int]*SimpleGraphEdge{},
		Attributes: map[string]string{}}
}

func (n SimpleGraphNode) Id() int {
	return n.id
}

// SimpleGraphEdge is the edge structure thats is stored in SimpleGraph and that
// is stores Nodes connected to it in a list
type SimpleGraphEdge struct {
	id int
	// First one is the input Node and second one is the output Node
	Nodes      []*SimpleGraphNode
	Attributes map[string]string
}

func NewSimpleGraphEdge() *SimpleGraphEdge {
	return &SimpleGraphEdge{Nodes: []*SimpleGraphNode{},
		Attributes: map[string]string{}}
}

func (e SimpleGraphEdge) Id() int {
	return e.id
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

func (g SimpleGraph) Edge(id int) *SimpleGraphEdge {
  return g.edges[id]
}
func (g SimpleGraph) Node(id int) *SimpleGraphNode {
  return g.nodes[id]
}

// AddNode adds a node to the SimpleGraph and returns it
func (g *SimpleGraph) AddNode() *SimpleGraphNode {
	n := NewSimpleGraphNode()

	id := g.lastNodeID + 1
	g.lastNodeID = id
	n.id = id

	g.nodes[id] = n
	return n
}

// AddEdge adds an edge to the SimpleGraph connected to fromNode and toNode
func (g *SimpleGraph) AddEdge(fromNode *SimpleGraphNode,
	toNode *SimpleGraphNode) *SimpleGraphEdge {

	id := g.lastEdgeID + 1
	g.lastEdgeID = id

	e := NewSimpleGraphEdge()
	e.id = id
	e.Nodes = append(e.Nodes, fromNode)
	e.Nodes = append(e.Nodes, toNode)

	fromNode.Edges[id] = e
	toNode.Edges[id] = e

	g.edges[id] = e
	return e
}

// RemoveNode removes n from SimpleGraph, if there are edges connected to that
// node they are also destroyed
func (g *SimpleGraph) RemoveNode(n *SimpleGraphNode) (err error) {

	delete(g.nodes, n.id)

	for _, e := range n.Edges {
		g.RemoveEdge(e)
	}

	return nil
}

// RemoveEdge removes the edge e from the SimpleGraph and deletes the link to it
// from the nodes it connects together
func (g *SimpleGraph) RemoveEdge(e *SimpleGraphEdge) (err error) {

	delete(g.edges, e.id)

	for _, n := range e.Nodes {
		delete(n.Edges, e.id)
	}

	return nil
}
