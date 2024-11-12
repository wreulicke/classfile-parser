package examples

import (
	"os"
	"testing"

	parser "github.com/wreulicke/classfile-parser"
)

func TestConstantPool(t *testing.T) {
	t.Parallel()

	f, err := os.Open("../testdata/classes/main/Test.class")
	if err != nil {
		t.Fatal(err)
	}
	cf, err := parser.New(f).Parse()
	if err != nil {
		t.Error(err)
	}
	for i, e := range cf.ConstantPool.Constants {
		if e == nil { // skip nil when last element is Long or Double
			continue
		}
		t.Logf("%d, %s", i, e.Name())
	}
}
