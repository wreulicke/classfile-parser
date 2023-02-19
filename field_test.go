package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFieldName(t *testing.T) {
	cf, err := parseFile("./testdata/classes/main/Test.class")
	assert.NoError(t, err)

	for _, e := range cf.Fields {
		_, err = e.Name(cf.ConstantPool)
		assert.NoError(t, err)
		_, err = e.Descriptor(cf.ConstantPool)
		assert.NoError(t, err)
	}
}
