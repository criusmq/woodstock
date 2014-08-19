package main

import (
	"bufio"
	"flag"
	"fmt"
	/* "github.com/codegangsta/negroni" */
	"encoding/json"
	"github.com/criusmq/woodstock/graph"
	"github.com/criusmq/woodstock/importer"
	"log"
	/* "net/http" */
	"os"
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

	s, err := json.MarshalIndent(g, "", " ")
	fmt.Printf("Thegraph \n ========\n%s\n%v", s, err)

	/* mux := http.NewServeMux() */
	/* mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) { */
	/* 	fmt.Fprintf(w, "Welcome to the home page!") */
	/* }) */

	/* n := negroni.Classic() */

	/* n.UseHandler(mux) */
	/* n.Run(":3000") */
}
