package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("./testdata/Test.class")
	if err != nil {
		t.Fatal(err)
	}
	p := New(f)
	c, err := p.Parse()
	if err != nil {
		t.Error(err)
	}
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")
	e.Encode(c)
	t.Error("test")
}
