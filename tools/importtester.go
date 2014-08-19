package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/criusmq/woodstock/importer"
	"github.com/criusmq/woodstock/graph"
	"log"
	"os"
  "encoding/json"
)

var inputFile = flag.String("infile", "simple.spept", "Input file path")

func main() {
	flag.Parse()

	fi, err := os.Open(*inputFile)

	if err != nil {
		log.Fatal(err)
		return
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(fi)
	v := importer.ImportPetriNet(r)
	g := graph.NewSimpleGraph()
	v.Graph(g)
fmt.Println("%v",v)
fmt.Println("%v",g)

b, err := json.Marshal(g)
if err != nil {
    fmt.Println("error:", err)
}
os.Stdout.Write(b)


  fmt.Printf("done.....................")

}
