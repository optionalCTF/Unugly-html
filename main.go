package main

import (
	"log"
	"os"

	"unugly/client"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("HTML No Ugly", "No more ugly HTML...")

	webAddr := parser.String("u", "url", &argparse.Options{Required: true, Help: "Supply a target URL to pull ugly HTML from.."})
	outPath := parser.String("o", "outfile", &argparse.Options{Required: false, Help: "Supply an output file"})
	err := parser.Parse(os.Args)

	if err != nil {
		log.Fatal(parser.Usage(err))
	}

	client.Pull(*webAddr, *outPath)
}
