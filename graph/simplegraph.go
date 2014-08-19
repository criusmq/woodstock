package graph

// SimpleGraphVertex is the vertex structure that is stored in SimpleGraph. It also
// link to the different edges which are connected to it.
type SimpleGraphVertex struct {
	id         int
	edges      map[int]*SimpleGraphEdge
	attributes map[string]interface{}
}

func (v SimpleGraphVertex) Id() int { return v.id }

func (v SimpleGraphVertex) Edges() map[int]Edge {
	// Cast isn't automatic in go for memory layouting so lets copy in a new map
	// this may not be the optimal way for nodes with lots of edges
	edges := make(map[int]Edge, len(v.edges))

	for key, value := range v.edges {
		edges[key] = value
	}
	return edges
}

func (v SimpleGraphVertex) Attributes() map[string]interface{} {
	return v.attributes
}

func newSimpleGraphVertex() *SimpleGraphVertex {
	return &SimpleGraphVertex{
		edges:      map[int]*SimpleGraphEdge{},
		attributes: map[string]interface{}{}}
}

// SimpleGraphEdge is the edge structure 
type SimpleGraphEdge struct {
	id         int
	vertices   []*SimpleGraphVertex
	attributes map[string]interface{}
}

func (e SimpleGraphEdge) Id() int {
	return e.id
}
func (e SimpleGraphEdge) Vertices() []Vertex {
	vertices := make([]Vertex, len(e.vertices))

	for key, value := range e.vertices {
		vertices[key] = value
	}

	return vertices
}

func (e SimpleGraphEdge) Attributes() map[string]interface{} {
	return e.attributes
}

func newSimpleGraphEdge() *SimpleGraphEdge {
	return &SimpleGraphEdge{
		vertices:   []*SimpleGraphVertex{},
		attributes: map[string]interface{}{}}
}

type SimpleGraph struct {
	vertices map[int]Vertex
	edges    map[int]Edge

	lastEdgeID int
	lastNodeID int
}

func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{
		lastEdgeID: -1,
		lastNodeID: -1,
		vertices:   map[int]Vertex{},
		edges:      map[int]Edge{}}
}

func (g SimpleGraph) Vertices() map[int]Vertex {
	// Cast isn't automatic in go for memory layouting so lets copy in a new map
	// this may not be the optimal way for nodes with lots of edges
	vertices := make(map[int]Vertex, len(g.vertices))

	for key, value := range g.vertices {
		vertices[key] = value
	}
	return vertices
}

func (g SimpleGraph) Edges() map[int]Edge {
	// Cast isn't automatic in go for memory layouting so lets copy in a new map
	// this may not be the optimal way for nodes with lots of edges
	edges := make(map[int]Edge, len(g.edges))

	for key, value := range g.edges {
		edges[key] = value
	}
	return edges
}

// AddVertex adds a node to the SimpleGraph and returns it
func (g *SimpleGraph) AddVertex() Vertex {
	v := *newSimpleGraphVertex()

	g.lastNodeID = g.lastNodeID + 1
	v.id = g.lastNodeID

	g.vertices[g.lastNodeID] = v
	return v
}

// AddEdge adds an edge to the SimpleGraph connected to fromNode and toNode
func (g *SimpleGraph) AddEdge(fromVertex Vertex, toVertex Vertex) Edge {
	e := *newSimpleGraphEdge()

	g.lastEdgeID = g.lastEdgeID + 1
	e.id = g.lastEdgeID

	fromSimpleVertex := fromVertex.(SimpleGraphVertex) 
	toSimpleVertex := toVertex.(SimpleGraphVertex)

	e.vertices = append(e.vertices, &fromSimpleVertex)
	e.vertices = append(e.vertices, &toSimpleVertex)

	g.edges[g.lastEdgeID] = e
	return e
}

// RemoveVertex removes Vertex v from SimpleGraph, if there are edges connected
// to that node they are also destroyed
func (g *SimpleGraph) RemoveVertex(v Vertex) {
	delete(g.Vertices(), v.Id())

	for _, e := range v.Edges() {
		g.RemoveEdge(e)
	}

}

// RemoveEdge removes the edge e from the SimpleGraph and deletes the link to it
// from the nodes it connects together
func (g *SimpleGraph) RemoveEdge(e Edge) {
	delete(g.edges, e.Id())

	for _, n := range e.Vertices() {
		delete(n.Edges(), e.Id())
	}
}

var _ Vertex = SimpleGraphVertex{}
var _ Edge = SimpleGraphEdge{}
var _ Graph = SimpleGraph{}
