package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func init() {
	parser := argparse.NewParser("HTML No Ugly", "No more ugly HTML...")

	webAddr := parser.String("u", "url", &argparse.Options{Required: true, Help: "Supply a target URL to pull ugly HTML from.."})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Println(parser.Usage(err))
	}
}

func main() {
	fmt.Println("reee")
}

func httpGet(url string) {

}
