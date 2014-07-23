package importer

import (
	"encoding/xml"
	"fmt"
	"github.com/criusmq/woodstock/graph"
	"io"
	"strings"
  "strconv"
)

// Snoopy
//  > NodeClasses "Places" "Transitions" "Coarse Place" "Coarse Transition"
//    > NodeClass
//      > Node
//        > Attribute
// > EdgeClasses "Edges" "Read Edges" "Inhibitor Edge" "Reset Edge" "Equal Edge"
//   > EdgeClass
//     > Edge
//        > Attribute

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
	Id         int         `xml:"id,attr"`
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
	g := graph.NewSimpleGraph()

	// Simple Map since node ids are gonna change
	nodes := map[int]*graph.SimpleGraphNode{}

	// for each node create a node
	for _, nc := range S.NodeClasses {
		for _, n := range nc.Nodes {
			node := g.AddNode()
			nodes[n.Id] = node
			fmt.Printf("NodeClass=%v, Node.id=%v, g.Node.Id=%v\n", nc.Name, n.Id, node.Id())
		}
	}

	// for each edge create an edge connecting the corresponding nodes
	for _, ec := range S.EdgeClasses {
		for _, e := range ec.Edges {
      
      multiplicity := 0 
			// collect the needed attributes
			for _, a := range e.Attributes {
				content := strings.Trim(a.Content, "\n\r ")

				switch a.Name {
				case "Multiplicity":
          i64multiplicity,_ := strconv.ParseInt(content,10,32)
          multiplicity = int(i64multiplicity)
				}

			}
			// Add the edge to the graph
			edge := g.AddEdge(nodes[e.Source], nodes[e.Target])
			fmt.Printf("EdgeClass=%v, Edge=%v -> %v(src=%v,dst=%v)\n", ec.Name, e.Id, edge.Id(), e.Source, e.Target)
      fmt.Printf("multiplicity=%v\n",multiplicity)
		}
	}
	return g
}
