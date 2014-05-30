package graph

type SimpleGraphNode struct {
	id          int
	edges       []*Edge
	linkedNodes []*Node
}

func (n SimpleGraphNode) Id() int              { return n.id }
func (n SimpleGraphNode) Edges() []*Edge       { return n.edges }
func (n SimpleGraphNode) LinkedNodes() []*Node { return n.linkedNodes }

type SimpleGraphEdge struct {
	id    int
	nodes []*Node
}

func (e SimpleGraphEdge) Id() int        { return e.id }
func (e SimpleGraphEdge) Nodes() []*Node { return e.nodes }

// SimpleGraph is a simple non directed graph which respects the Graph interface
type SimpleGraph struct {
	nodes map[int]*SimpleGraphNode
	edges map[int]*SimpleGraphEdge
}

func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{}
}

func (g SimpleGraph) addNode(node *Node) {
}

func (g SimpleGraph) addEdge(edge *Edge, fromNode *Node, toNode *Node) {
}

func (g SimpleGraph) removeNode(node *Node) {
}
func (g SimpleGraph) removeEdge(edge *Edge) {
}

func (g SimpleGraph) Node(id int) *Node { return nil }
func (g SimpleGraph) Edge(id int) *Edge { return nil }

// Verify interface satisfaction
var _ Node = SimpleGraphNode{}
var _ Edge = SimpleGraphEdge{}
var _ Graph = SimpleGraph{}
