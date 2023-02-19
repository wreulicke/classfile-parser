package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadMethodName(t *testing.T) {
	cf, err := parseFile("./testdata/classes/main/Test.class")
	assert.NoError(t, err)

	for _, e := range cf.Methods {
		_, err = e.Name(cf.ConstantPool)
		assert.NoError(t, err)
		_, err = e.Descriptor(cf.ConstantPool)
		assert.NoError(t, err)
	}
}

func TestAnnotationDefault(t *testing.T) {
	cf, err := parseFile("./testdata/classes/main/Annot.class")
	assert.NoError(t, err)

	m := cf.Methods[0]

	name, _ := m.Name(cf.ConstantPool)
	assert.Equal(t, "value", name)
	attr := m.AnnotationDefault()
	assert.NotNil(t, attr)

	index := attr.DefaultValue.(*ElementValueConstValue).ConstValueIndex
	utf8, err := cf.ConstantPool.GetConstantUtf8(index)
	assert.NoError(t, err)
	assert.Equal(t, "default", utf8.String())
}
