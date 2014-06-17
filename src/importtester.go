package main

import (
	"bufio"
	"flag"
	"log"
	"os"
  "fmt"
	"woodstock/importer"
)

var inputFile = flag.String("infile", "enwiki-latest-pages-articles.xml", "Input file path")

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

  v:= importer.ImportPetriNet(r)
	fmt.Printf("Snoopy=%+v", *v)
  
}