package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parseFile(path string) (*Classfile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	p := New(f)
	return p.Parse()
}

func TestParse(t *testing.T) {
	filepath.Walk("./testdata", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".class") {
			return nil
		}
		t.Run(path, func(t *testing.T) {
			fmt.Printf("============================== %s: start ===============================\n", path)
			f, err := os.Open(path)
			if err != nil {
				t.Fatal(err)
			}
			p := New(f)
			cf, err := p.Parse()
			if err != nil {
				t.Error(path, err)
			}
			if !strings.HasSuffix(path, "module-info.class") {
				_, err = cf.ThisClassName()
				assert.NoError(t, err)
				_, err = cf.SuperClassName()
				assert.NoError(t, err)
			}
			fmt.Printf("============================== %s: end ===============================\n", path)
		})
		return nil
	})
}
