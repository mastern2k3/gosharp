package gosharp

import (
	"reflect"
	"testing"
)

const (
	Hello = 123
)

func TestExtractConsts(t *testing.T) {

	expected := []ConstDef{
		ConstDef{
			name: "Hello",
		},
	}

	defs, err := ExtractConsts(`
	package lol
	const (
		Hello = 123
	)`)

	if err != nil {
		t.Fatalf("error while calling ConstNames: %s", err)
	}

	if !reflect.DeepEqual(defs, expected) {
		t.Fatalf("definitions found not matching expected, expected %+v, found %+v", expected, defs)
	}
}
