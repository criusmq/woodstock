// The package graph is a package comprising of multiple implementation of graphs
// all based on the graph interface
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

// Special type of Edge which is used in directed graph implementation
type DirectedEdge interface {
	Id() int
	Nodes() [](*Node)
	From() *Node
	To() *Node
}

// Graph interface is an interface to permit dependency injection of different
// graph types in the same application. It's a structure with node/vertex and 
// edges/arcs there is no dinstinction in the fromNode and toNode if it is a
// normal Edge but there is a distinction in with a DirectedEdge 
//(Maybe a special interface with the name DirectedGraph should be created).
type Graph interface {
	addNode(node *Node)
	addEdge(edge *Edge, fromNode *Node, toNode *Node)

	removeNode(node *Node)
	removeEdge(edge *Edge)

	Node(id int) *Node
	Edge(id int) *Edge
}
