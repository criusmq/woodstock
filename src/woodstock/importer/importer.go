package importer

import (
	"io"
	"woodstock/graph"
	"encoding/xml"
)

// Snoopy
//  > NodeClasses "Places" "Transitions" "Coarse Place" "Coarse Transition"
//   > NodeClass
//     > Node
// > EdgeClasses "Edges" "Read Edges" "Inhibitor Edge" "Reset Edge" "Equal Edge"
//   > EdgeClass
//     > Edge

type Snoopy struct {
	NodeClasses []NodeClass `xml:"nodeclasses>nodeclass"`
	EdgeClasses []EdgeClass `xml:"edgeclasses>edgeclass"`

	Version  string `xml:"version,attr"`
	Revision string `xml:"revision,attr"`
}

type NodeClass struct {
	Nodes []Node `xml:"node"`
	Name  string `xml:"name,attr"`
}

type Node struct {
	Id string `xml:"id,attr"`
  Attributes []Attribute `xml:"attribute"`
}

type EdgeClass struct {
	Edges []Edge `xml:"edge"`
	Name  string `xml:"name,attr"`
}
type Edge struct {
	Id string `xml:"id,attr"`
  Attributes []Attribute `xml:"attribute"`
}
type Attribute struct{
  Name string `xml:"name,attr"`
}

// ImportPetriNet imports a petrinet into a graph "g"
func ImportPetriNet(r io.Reader) *Snoopy{
	v := Snoopy{}
	xml.NewDecoder(r).Decode(&v)

  return &v
}

func (S Snoopy) graph() *graph.SimpleGraph{
  return graph.NewSimpleGraph()
}