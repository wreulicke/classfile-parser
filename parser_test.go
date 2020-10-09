package main

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("./Test.class")
	if err != nil {
		t.Fatal(err)
	}
	p := New(f)
	_, err = p.Parse()
	if err != nil {
		t.Error(err)
	}
}