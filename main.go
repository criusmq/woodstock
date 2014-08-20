package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/codegangsta/negroni"
	"encoding/json"
	"github.com/criusmq/woodstock/graph"
	"github.com/criusmq/woodstock/importer"
  "github.com/gorilla/mux"
	"log"
	"net/http"
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
	
  // graph is now generated
  v.Graph(g)

	s, err := json.MarshalIndent(g, "", "\t")
  
	router := mux.NewRouter()
  router.HandleFunc("/graph.json",func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s",s)
	})

	n := negroni.Classic()

	n.UseHandler(router)
	n.Run(":3000")
}
