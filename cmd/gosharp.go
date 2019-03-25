package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mastern2k3/gosharp"
)

var (
	goFile = flag.String("i", "", "the input go file to parse")
)

func main() {

	s, err := os.Stat(*goFile)
	if err != nil {
		panic(err)
	}

	if s.IsDir() {
		panic(fmt.Sprintf("%s is a directory, expected go source file", *goFile))
	}

	gosharp.ExtractConsts("")
}
