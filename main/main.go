package main

import (
	"flag"
	"fmt"
	link "github.com/mdurokov/GoLinkParser"

	"log"
	"os"
)

type Link struct {
	Href string
	Text string
}

func main() {

	fileName := flag.String("file", "ex3.html", "HTML file to parse and get links from it")
	flag.Parse()

	htmlFile, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	links, err := link.Parse(htmlFile)

	fmt.Printf("Links: %+v", links)

}
