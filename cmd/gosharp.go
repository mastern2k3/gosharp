package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/mastern2k3/gosharp"
)

var (
	input  = flag.String("i", "", "the input go file to parse")
	output = flag.String("o", "", "the output csharp file to generate")
)

var (
	csharpTemplate = template.Must(template.New("").Parse(`
{{- /**/ -}}
static class Constants {
{{range .}}    public const int {{.Name}} = 3000;
{{end -}}
}
`))
)

func main() {

	flag.Parse()

	if *input == "" || *output == "" {
		panic("must provide input and output files via -i and -o")
	}

	s, err := os.Stat(*input)
	if err != nil {
		panic(fmt.Sprintf("error while inspecting input: %s", err))
	}

	if s.IsDir() {
		panic(fmt.Sprintf("%s is a directory, expected go source file", *input))
	}

	f, err := os.Open(*input)
	if err != nil {
		panic(fmt.Sprintf("error while opening file: %s", err))
	}

	defs, err := gosharp.ExtractConstsReader(bufio.NewReader(f), path.Base(*input))
	if err != nil {
		panic(fmt.Sprintf("error while parsing file, make sure it compiles before generating: %s", err))
	}

	if err := csharpTemplate.Execute(os.Stdout, defs); err != nil {
		panic(fmt.Sprintf("error while producing output: %s", err))
	}
}
