package graph

type SimpleGraphNode struct {
	id    int
	edges map[int]*SimpleGraphEdge
}

func NewSimpleGraphNode() *SimpleGraphNode {
	return &SimpleGraphNode{edges: map[int]*SimpleGraphEdge{}}
}

type SimpleGraphEdge struct {
	id int
	// First one is the input Node and second one is the output Node
	nodes []*SimpleGraphNode
}

func NewSimpleGraphEdge() *SimpleGraphEdge {
	return &SimpleGraphEdge{nodes: []*SimpleGraphNode{}}
}

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

func (g SimpleGraph) addNode() *SimpleGraphNode {
	n := NewSimpleGraphNode()

	id := g.lastNodeID + 1
	n.id = id
	g.lastNodeID = id

	g.nodes[id] = n
	return n
}

func (g SimpleGraph) addEdge(fromNode *SimpleGraphNode,
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

func (g SimpleGraph) removeNode(n *SimpleGraphNode) (err error) {

	delete(g.nodes, n.id)

	for _, e := range n.edges {
		g.removeEdge(e)
	}

	return nil
}

func (g SimpleGraph) removeEdge(e *SimpleGraphEdge) (err error) {

	delete(g.edges, e.id)

	for _, n := range e.nodes {
		delete(n.edges, e.id)
	}

	return nil
}
