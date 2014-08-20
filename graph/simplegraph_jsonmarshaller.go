package graph

import (
	"encoding/json"
	"strconv"
)

type jsonVertex struct {
	Id         int                    `json:"id"`
	Edges      []string               `json:"edges"`
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

func newJsonVertex(v SimpleGraphVertex) *jsonVertex {
	jv := &jsonVertex{
		Id:         v.Id(),
		Attributes: v.Attributes()}

	edges := v.Edges()
	jv.Edges = make([]string, 0, len(edges))

	for _, edge := range v.edges {
		jv.Edges = append(jv.Edges, strconv.Itoa(edge.Id()))
	}

	return jv
}

func newJsonEdge(e SimpleGraphEdge) *jsonEdge {
	je := &jsonEdge{Id: e.Id(), Attributes: e.Attributes()}

	je.Vertices = []string{}

	for _, vertex := range e.Vertices() {
		je.Vertices = append(je.Vertices, strconv.Itoa(vertex.Id()))
	}
	return je
}

func newJsonGraph(g SimpleGraph) *jsonGraph {
	jg := &jsonGraph{
		Vertices: map[string]*jsonVertex{},
		Edges:    map[string]*jsonEdge{},
	}

	for key, vertex := range g.Vertices() {
		jg.Vertices[strconv.Itoa(key)] = newJsonVertex(*vertex)
	}
	for key, edge := range g.Edges() {
		jg.Edges[strconv.Itoa(key)] = newJsonEdge(*edge)
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
