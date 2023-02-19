package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnnotation(t *testing.T) {
	cf, err := parseFile("./testdata/classes/main/Test.class")
	assert.NoError(t, err)

	m := cf.Methods[1] // expected main

	name, _ := m.Name(cf.ConstantPool)
	assert.Equal(t, "main", name)

	desc, _ := m.Descriptor(cf.ConstantPool)
	assert.Equal(t, "([Ljava/lang/String;)V", desc)
	attr := findAttribute(m.Attributes, "RuntimeVisibleAnnotations").(*AttributeRuntimeVisibleAnnotations)
	annot := attr.Annotations[1]
	typ, _ := annot.Type(cf.ConstantPool)
	assert.Equal(t, "Lmain/Annot;", typ)
	assert.Empty(t, annot.ElementValuePairs)
}
