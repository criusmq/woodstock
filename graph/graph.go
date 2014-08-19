// The package graph is a package comprising of multiple implementation of
// graphs all based on the graph interface
package graph

// Vertex is the basic element of a graph
type Vertex interface {
  // Id to be mappable
	Id() int
	Edges() map[int]Edge
	Attributes() map[string]interface{}
}

// Edges are the basic element of a Graph that links Nodes together
type Edge interface {
  // Id to be mappable
	Id() int
  // On a directed graph the first Vertex in Vertices is the input vertex
  // the second one is the ouput Vertex
	Vertices() []Vertex
	Attributes() map[string]interface{}
}

// The Graph container interface, it contains associated edges and nodes with
// their relations
type Graph interface {
	// This method permits to retrieve all vertices with their associated id stored
	// in the graph
	Vertices() map[int]Vertex
	// This method permits to retrieve all edges with their associated id stored
	// in the graph
	Edges() map[int]Edge

	AddVertex() Vertex
	AddEdge(fromVertex Vertex, toVertex Vertex) Edge

	RemoveVertex(v Vertex)
	RemoveEdge(e Edge)
}
