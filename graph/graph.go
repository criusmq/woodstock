package graph


type Vertex interface{
  Id() int
  Edges() map[int]Edge
  Attributes() map[string]interface{}
}
type Edge interface{
  Id() int
  Vertices() []Vertex
  Attributes() map[string]interface{}
}
type Graph interface{
  Vertices() map[int]Vertex
  Edges() map[int]Edge
  AddVertex() Vertex
  AddEdge(fromVertex Vertex, toVertex Vertex) Edge

  RemoveVertex(v Vertex)
  RemoveEdge(e Edge)
}