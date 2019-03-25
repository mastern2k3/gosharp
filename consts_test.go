package gosharp

import (
	"bufio"
	"go/token"
	"os"
	"reflect"
	"testing"
)

func TestExtractConstsString(t *testing.T) {

	expected := []ConstDef{
		ConstDef{
			Name:  "Hello",
			Value: "123",
			Type:  token.INT,
		},
		ConstDef{
			Name:  "World",
			Value: "456",
			Type:  token.INT,
		},
	}

	defs, err := ExtractConstsString(`
	package test
	
	const (
		Hello = 123
		World = 456
	)`)

	if err != nil {
		t.Fatalf("error while calling ConstNames: %s", err)
	}

	if !reflect.DeepEqual(defs, expected) {
		t.Fatalf("definitions found not matching expected, expected %+v, found %+v", expected, defs)
	}
}

/*
const (
	Shalom = 123
	Olam   = 456
)
*/
func TestExtractConstsReader(t *testing.T) {

	expected := []ConstDef{
		ConstDef{
			Name:  "Shalom",
			Value: "123",
			Type:  token.INT,
		},
		ConstDef{
			Name:  "Olam",
			Value: "456",
			Type:  token.INT,
		},
	}

	f, err := os.Open("./testdata/testsrc.go")

	if err != nil {
		t.Fatalf("error while opening test file: %s", err)
	}

	defs, err := ExtractConstsReader(bufio.NewReader(f), "testsrc.go")

	if err != nil {
		t.Fatalf("error while calling ConstNames: %s", err)
	}

	if !reflect.DeepEqual(defs, expected) {
		t.Fatalf("definitions found not matching expected, expected %+v, found %+v", expected, defs)
	}
}
