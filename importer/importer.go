package importer

import (
  "encoding/xml"
  "github.com/criusmq/woodstock/graph"
  "io"
  "strconv"
  "strings"
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
  Attributes []NodeAttribute `xml:"attribute"`
  Graphics   NodeGraphic    `xml:"graphics>graphic"`
}

type EdgeClass struct {
  Edges []Edge `xml:"edge"`
  Name  string `xml:"name,attr"`
}
type Edge struct {
  Id         int         `xml:"id,attr"`
  Source     int         `xml:"source,attr"`
  Target     int         `xml:"target,attr"`
  Attributes []EdgeAttribute `xml:"attribute"`
  Graphics EdgeGraphic `xml:"graphics>graphic"`
}
type NodeAttribute struct {
  Id       int     `xml:"id,attr"`
  Name     string  `xml:"name,attr"`
  Content  string  `xml:",chardata"`
  Graphics NodeAttributeGraphic `xml:"graphics>graphic"`
}
type EdgeAttribute struct {
  Id       int     `xml:"id,attr"`
  Name     string  `xml:"name,attr"`
  Content  string  `xml:",chardata"`
}

type NodeGraphic struct {
  Id   int     `xml:"id,attr`
  X    float64 `xml:"x,attr"`
  Y    float64 `xml:"y,attr"`
  W    float64 `xml:"w,attr"`
  H    float64 `xml:"h,attr"`
  Net  int     `xml:"net,attr"`
  Show int     `xml:"show,attr"`
}
type NodeAttributeGraphic struct {
  Id   int     `xml:"id,attr"`
  Net  int     `xml:"net,attr"`
  Show int     `xml:"show,attr"`
  Yoff float64 `xml:"yoff,attr"`
  Xoff float64 `xml:"xoff,attr"`
  X    float64 `xml:"x,attr"`
  Y    float64 `xml:"y,attr"`
}
type EdgeGraphic struct {
  Id   int     `xml:"id,attr"`
  Net  int     `xml:"net,attr"`
  Show int     `xml:"show,attr"`
  Points []Point `xml:"points>point"`
}
type Point struct {
  X float64 `xml:"x,attr"`
  Y float64 `xml:"y,attr"`
}

// ImportPetriNet imports a snoopy spept  file into a usable structure in woodstock
func ImportPetriNet(r io.Reader) *Snoopy {
  v := Snoopy{}
  xml.NewDecoder(r).Decode(&v)

  return &v
}

// Shall convert the Snoopy structure S to a new graph
func (S Snoopy) Graph(g *graph.SimpleGraph) {

  // Simple Map since node ids are gonna change
  nodes := map[int]*graph.SimpleGraphVertex{}

  // for each node create a node
  for _, nc := range S.NodeClasses {
    for _, n := range nc.Nodes {
      node := g.AddVertex()
      nodes[n.Id] = node

      attr := node.Attributes()
      attr["type"] = nc.Name
      
      for _, attribute := range n.Attributes{
        switch attribute.Name{
        case "Name":
          attr["name"]=attribute.Content
          attr["namePosition"] = attribute.Graphics
        }
      }
      attr["graphics"] = n.Graphics
      //			fmt.Printf("Node = %p %v\n",g.Node(node.Id()), node)
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
          multiplicity, _ = strconv.Atoi(content)
        }

      }
      // Add the edge to the graph
      edge := g.AddEdge(nodes[e.Source], nodes[e.Target])

      attr := edge.Attributes()
      attr["multiplicity"] = strconv.Itoa(multiplicity)
      attr["type"] = ec.Name

      attr["graphics"] = e.Graphics 
    }
  }
}
