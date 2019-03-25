package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/token"
	"io"
	"os"
	"path"
	"text/template"

	"github.com/mastern2k3/gosharp"
)

var (
	output = flag.String("o", "", "the output csharp file to generate (default to stdout)")

	classname = flag.String("classname", "Constants", "the class name to use when generating the csharp file")
)

var csharpTemplate = template.Must(template.New("").
	Funcs(template.FuncMap{
		"typeOf": typeOf,
	}).Parse(`
{{- /**/ -}}
static class {{.ClassName}} {
{{range .Defs}}    public const {{typeOf .}} {{.Name}} = {{.Value}};
{{end -}}
}
`))

type templateData struct {
	ClassName string
	Defs      []gosharp.ConstDef
}

func typeOf(def gosharp.ConstDef) (string, error) {
	switch def.Type {
	case token.INT:
		return "int", nil
	default:
		return "", fmt.Errorf("could not determine type of definition %+v", def)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of gosharp:\n")
	fmt.Fprintf(os.Stderr, "\tgosharp [flags] file_with_consts.go\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		flag.Usage()
		os.Exit(2)
	}

	input := args[0]

	s, err := os.Stat(input)
	if err != nil {
		panic(fmt.Sprintf("error while inspecting input: %s", err))
	}

	if s.IsDir() {
		panic(fmt.Sprintf("%s is a directory, expected go source file", input))
	}

	f, err := os.Open(input)
	if err != nil {
		panic(fmt.Sprintf("error while opening file: %s", err))
	}

	defs, err := gosharp.ExtractConstsReader(bufio.NewReader(f), path.Base(input))
	if err != nil {
		panic(fmt.Sprintf("error while parsing file, make sure it compiles before generating: %s", err))
	}

	var writer io.Writer

	if *output == "" {
		writer = os.Stdout
	} else {
		f, err := os.Create(*output)
		if err != nil {
			panic(fmt.Sprintf("error while creating output file: %s", err))
		}

		defer f.Close()

		writer = f
	}

	data := templateData{
		Defs:      defs,
		ClassName: *classname,
	}

	if err := csharpTemplate.Execute(writer, data); err != nil {
		panic(fmt.Sprintf("error while producing output: %s", err))
	}
}
