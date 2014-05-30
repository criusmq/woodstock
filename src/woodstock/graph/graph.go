package graph

type Node interface {
	Id() int
	Edges() [](*Edge)
	LinkedNodes() [](*Node)
}

type Edge interface {
	Id() int
	Nodes() [](*Node)
}

type DirectedEdge interface {
	Id() int
	Nodes() [](*Node)
	From() *Node
	To() *Node
}

type Graph interface {
	addNode(node *Node)
	addEdge(edge *Edge, fromNode *Node, toNode *Node)

	removeNode(node *Node)
	removeEdge(edge *Edge)

	Node(id int) *Node
	Edge(id int) *Edge
}

