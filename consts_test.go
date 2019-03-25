package gosharp

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func TestExtractConstsString(t *testing.T) {

	expected := []ConstDef{
		ConstDef{
			name: "Hello",
		},
		ConstDef{
			name: "World",
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
			name: "Shalom",
		},
		ConstDef{
			name: "Olam",
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
