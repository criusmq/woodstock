package importer

import (
	"encoding/xml"
	"github.com/criusmq/woodstock/graph"
	"io"
  "fmt"
  "strings"
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
	Id         string      `xml:"id,attr"`
	Attributes []Attribute `xml:"attribute"`
}

type EdgeClass struct {
	Edges []Edge `xml:"edge"`
	Name  string `xml:"name,attr"`
}
type Edge struct {
	Id         int         `xml:"id,attr"`
	Source     int         `xml:"source,attr"`
	Target     int         `xml:"target,attr"`
	Attributes []Attribute `xml:"attribute"`
}
type Attribute struct {
	Id      int    `xml:"id,attr"`
	Name    string `xml:"name,attr"`
	Content string `xml:",chardata"`
}

// ImportPetriNet imports a snoopy spept  file into a usable structure in woodstock
func ImportPetriNet(r io.Reader) *Snoopy {
	v := Snoopy{}
	xml.NewDecoder(r).Decode(&v)

	return &v
}

// Shall convert the Snoopy structure S to a new graph
func (S Snoopy) Graph() *graph.SimpleGraph {
  g:= graph.NewSimpleGraph()

  // for each node create a node
  for _,nc := range S.NodeClasses {
    for _,n := range nc.Nodes {
      fmt.Printf("NodeClass=%v, Node=%v\n" , nc.Name, n.Id)
      // collect the needed attributes
    }
  }

  // for each edge create an edge connecting the corresponding nodes
  for _,ec := range S.EdgeClasses {
    for _,e := range ec.Edges {
      fmt.Printf("EdgeClass=%v, Edge=%v(src=%v,dst=%v)\n" , ec.Name, e.Id,e.Source,e.Target)

      // collect the needed attributes
      for _,a := range e.Attributes{
        content:= strings.Trim(a.Content,"\n\r ")
        
        switch a.Name{
          case "Multiplicity": fmt.Printf("Attribute=%v, Content=%v\n",a.Name,content)
        }
      }
    }
  }
	return g
}
