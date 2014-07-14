package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/criusmq/woodstock/importer"
	"log"
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
  
  v.Graph()
	//fmt.Printf("Snoopy=%+v", *v)
  fmt.Printf("done.....................")

}
