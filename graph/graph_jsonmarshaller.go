package graph

import (
	"encoding/json"
	"strconv"
)

type jsonVertex struct {
	Id         int                    `json:"id"`
	Edges      map[string]string      `json:"edges"`
	Attributes map[string]interface{} `json:"attributes"`
}
type jsonEdge struct {
	Id         int                    `json:"id"`
	Vertices   []string               `json:"vertices"`
	Attributes map[string]interface{} `json:"attributes"`
}

type jsonGraph struct {
	Vertices map[string]*jsonVertex `json:"vertices"`
	Edges    map[string]*jsonEdge   `json:"edges"`
}

func newJsonVertex(v Vertex) *jsonVertex {
	jv := &jsonVertex{
		Id:         v.Id(),
		Attributes: v.Attributes()}

	jv.Edges = map[string]string{}

	for key, edge := range v.Edges() {
		jv.Edges[strconv.Itoa(key)] = strconv.Itoa(edge.Id())
	}

	return jv
}

func newJsonEdge(e Edge) *jsonEdge {
	je := &jsonEdge{Id: e.Id(), Attributes: e.Attributes()}

	je.Vertices = []string{}

	for _, vertex := range e.Vertices() {
		je.Vertices = append(je.Vertices, strconv.Itoa(vertex.Id()))
	}
	return je
}

func newJsonGraph(g Graph) *jsonGraph {
	jg := &jsonGraph{
		Vertices: map[string]*jsonVertex{},
		Edges:    map[string]*jsonEdge{},
	}

	for key, vertex := range g.Vertices() {
		jg.Vertices[strconv.Itoa(key)] = newJsonVertex(vertex)
	}
	for key, edge := range g.Edges() {
		jg.Edges[strconv.Itoa(key)] = newJsonEdge(edge)
	}
	return jg
}

func (e SimpleGraphEdge) MarshalJSON() ([]byte, error) {
	return json.Marshal(newJsonEdge(e))
}
func (v SimpleGraphVertex) MarshalJSON() ([]byte, error) {
	return json.Marshal(newJsonVertex(v))
}
func (g SimpleGraph) MarshalJSON() ([]byte, error) {
	return json.Marshal(newJsonGraph(g))
}
