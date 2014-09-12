package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/criusmq/woodstock/graph"
	"github.com/criusmq/woodstock/importer"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
  "io"
)

var inputFile = flag.String("infile", "simple.spept", "Input file path")

func importGraph(r io.Reader) *graph.SimpleGraph{
  snoopyGraph := importer.ImportPetriNet(r)

  g := graph.NewSimpleGraph()

  snoopyGraph.Graph(g)

  return g
}

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


  var workGraph graph.SimpleGraph
	r := bufio.NewReader(fi)
  workGraph = *importGraph(r)


	router := mux.NewRouter()
	router.HandleFunc("/graph.json", func(w http.ResponseWriter, req *http.Request) {
    s, err := json.MarshalIndent(workGraph, "", "\t")

    if err == nil { return }
		
    fmt.Fprintf(w, "%s", s)
	})

	router.HandleFunc("/upload", func(w http.ResponseWriter, req *http.Request) {
		file, header, err := req.FormFile("graph")

		defer file.Close()

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
    
    workGraph = *importGraph(r)
    s, err := json.MarshalIndent(workGraph, "", "\t")
    fmt.Printf("%s", s)
    
		fmt.Fprintf(w, "File uploaded successfully : ")
		fmt.Fprintf(w, header.Filename)
	})

	n := negroni.Classic()

	n.UseHandler(router)
  var port = os.Getenv("PORT")
  if port == "" { port = "3000" }
	n.Run(":"+ port)
}
