package graph

// SimpleGraphVertex is the vertex structure that is stored in SimpleGraph. It also
// link to the different edges which are connected to it.
type SimpleGraphVertex struct {
	id         int
	edges      map[int]*SimpleGraphEdge
	attributes map[string]interface{}
}

func (v SimpleGraphVertex) Id() int                            { return v.id }
func (v SimpleGraphVertex) Edges() map[int]*SimpleGraphEdge    { return v.edges }
func (v SimpleGraphVertex) Attributes() map[string]interface{} { return v.attributes }

func newSimpleGraphVertex() *SimpleGraphVertex {
	return &SimpleGraphVertex{
		id:         -1,
		edges:      map[int]*SimpleGraphEdge{},
		attributes: map[string]interface{}{}}
}

// SimpleGraphEdge is the edge structure
type SimpleGraphEdge struct {
	id         int
	vertices   []*SimpleGraphVertex
	attributes map[string]interface{}
}

func (e SimpleGraphEdge) Id() int                            { return e.id }
func (e SimpleGraphEdge) Vertices() []*SimpleGraphVertex     { return e.vertices }
func (e SimpleGraphEdge) Attributes() map[string]interface{} { return e.attributes }

func newSimpleGraphEdge() *SimpleGraphEdge {
	return &SimpleGraphEdge{
		id:         -1,
		vertices:   []*SimpleGraphVertex{},
		attributes: map[string]interface{}{}}
}

type SimpleGraph struct {
	vertices map[int]*SimpleGraphVertex
	edges    map[int]*SimpleGraphEdge

	lastEdgeID int
	lastNodeID int
}

func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{
		lastEdgeID: -1,
		lastNodeID: -1,
		vertices:   map[int]*SimpleGraphVertex{},
		edges:      map[int]*SimpleGraphEdge{}}
}

func (g SimpleGraph) Vertices() map[int]*SimpleGraphVertex { return g.vertices }
func (g SimpleGraph) Edges() map[int]*SimpleGraphEdge      { return g.edges }

// AddVertex adds a node to the SimpleGraph and returns it
func (g *SimpleGraph) AddVertex() *SimpleGraphVertex {
	v := newSimpleGraphVertex()

	g.lastNodeID = g.lastNodeID + 1
	v.id = g.lastNodeID

	g.vertices[g.lastNodeID] = v
	return v
}

// AddEdge adds an edge to the SimpleGraph connected to fromNode and toNode
func (g *SimpleGraph) AddEdge(fromVertex *SimpleGraphVertex, toVertex *SimpleGraphVertex) *SimpleGraphEdge {
	e := newSimpleGraphEdge()

	g.lastEdgeID = g.lastEdgeID + 1
	e.id = g.lastEdgeID

	e.vertices = append(e.vertices, fromVertex)
	e.vertices = append(e.vertices, toVertex)

	fEdges := fromVertex.Edges()
	tEdges := toVertex.Edges()

	fEdges[e.id] = e
	tEdges[e.id] = e

	g.edges[g.lastEdgeID] = e
	return e
}

// RemoveVertex removes Vertex v from SimpleGraph, if there are edges connected
// to that node they are also destroyed
func (g *SimpleGraph) RemoveVertex(v *SimpleGraphVertex) {
	delete(g.Vertices(), v.Id())

	for _, e := range v.Edges() {
		g.RemoveEdge(e)
	}

}

// RemoveEdge removes the edge e from the SimpleGraph and deletes the link to it
// from the nodes it connects together
func (g *SimpleGraph) RemoveEdge(e *SimpleGraphEdge) {
	delete(g.edges, e.Id())

	for _, n := range e.Vertices() {
		delete(n.Edges(), e.Id())
	}
}

//var _ Vertex = SimpleGraphVertex{}
//var _ Edge = SimpleGraphEdge{}
//var _ Graph = SimpleGraph{}
