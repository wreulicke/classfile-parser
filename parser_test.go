package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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
			_, err = p.Parse()
			if err != nil {
				t.Error(path, err)
			}
			fmt.Printf("============================== %s: end ===============================\n", path)
		})
		return nil
	})
}
