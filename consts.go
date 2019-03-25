package gosharp

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
)

func ExtractConstsString(src string) ([]ConstDef, error) {
	return ExtractConsts(src, "")
}

func ExtractConstsReader(r io.Reader, filename string) ([]ConstDef, error) {
	return ExtractConsts(r, filename)
}

func ExtractConsts(src interface{}, filename string) ([]ConstDef, error) {

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filename, src, parser.AllErrors)

	if err != nil {
		return nil, err
	}

	v := constLocator{
		extractions: &[]ConstDef{},
	}

	ast.Walk(v, f)

	return *v.extractions, nil
}

type ConstDef struct {
	Name string
}

type constLocator struct {
	depth       int
	extractions *[]ConstDef
}

func (v constLocator) Visit(n ast.Node) ast.Visitor {

	if n == nil {
		return nil
	}

	if d, is := n.(*ast.GenDecl); is {
		if d.Tok == token.CONST {
			return valueLocator{depth: v.depth + 1, extractions: v.extractions}
		}
	}

	v.depth++

	return v
}

type valueLocator struct {
	depth       int
	extractions *[]ConstDef
}

func (v valueLocator) Visit(n ast.Node) ast.Visitor {

	if n == nil {
		return nil
	}

	if _, is := n.(*ast.ValueSpec); is {
		return &valueCapture{
			depth:       int(v.depth) + 1,
			extractions: v.extractions,
		}
	}

	v.depth++

	return v
}

type valueCapture struct {
	depth       int
	extractions *[]ConstDef
	name        string
	value       string
}

func (v *valueCapture) Visit(n ast.Node) ast.Visitor {

	if n == nil {
		return nil
	}

	if ent, is := n.(*ast.Ident); is {
		v.name = ent.Name
	} else if lit, is := n.(*ast.BasicLit); is {
		v.value = lit.Value
		ext := append(*v.extractions, ConstDef{
			Name: v.name,
		})
		*v.extractions = ext
	}

	return v
}
