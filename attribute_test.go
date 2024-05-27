package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findAttribute(attrs []Attribute, name string) Attribute {
	for _, e := range attrs {
		if e.Name() == name {
			return e
		}
	}
	return nil
}

func TestParseCode(t *testing.T) {
	t.Parallel()
	cf, err := parseFile("./testdata/classes/main/Test.class")
	assert.NoError(t, err)

	for _, e := range cf.Methods {
		attr := findAttribute(e.Attributes, "Code")
		if code, ok := attr.(*AttributeCode); ok {
			_, err := code.ParseCode()
			assert.NoError(t, err)
		}
	}
}
