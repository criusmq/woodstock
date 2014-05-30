package graph

type TransactionGraph struct {
}

func NewTransactionGraph() *TransactionGraph {
	return &TransactionGraph{}
}

func (g *TransactionGraph) addNode(node *Node)                               {}
func (g *TransactionGraph) addEdge(edge *Edge, fromNode *Node, toNode *Node) {}

func (g *TransactionGraph) removeNode(node *Node) {}
func (g *TransactionGraph) removeEdge(edge *Edge) {}

func (g *TransactionGraph) Node(id int) *Node { return nil }
func (g *TransactionGraph) Edge(id int) *Edge { return nil }
