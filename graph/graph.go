// The package graph is a package comprising of multiple implementation of graphs
// all based on the graph interface
package graph

/* // Node is a basic element of a graph */
/* type Node interface { */
/* 	ID() int */
/* 	Edges() [](*Edge) */
/* 	LinkedNodes() [](*Node) */
/* } */

/* // Edge is what connects nodes to other nodes in a graph */
/* type Edge interface { */
/* 	ID() int */
/* 	Nodes() [](*Node) */
/* } */

/* // DirectedEdge is a special type of Edge which is used in directed graph implementation */
/* type DirectedEdge interface { */
/* 	ID() int */
/* 	Nodes() [](*Node) */
/* 	From() *Node */
/* 	To() *Node */
/* } */

/* // Graph interface is an interface to permit dependency injection of different */
/* // graph types in the same application. It's a structure with node/vertex and */
/* // edges/arcs there is no dinstinction in the fromNode and toNode if it is a */
/* // normal Edge but there is a distinction in with a DirectedEdge */
/* //(Maybe a special interface with the name DirectedGraph should be created). */
/* type Graph interface { */
/* 	addNode() Node */
/* 	addEdge(fromNode *Node, toNode *Node) Edge */

/* 	removeNode(node *Node) */
/* 	removeEdge(edge *Edge) */

/* 	Node(id int) *Node */
/* 	Edge(id int) *Edge */
/* } */
