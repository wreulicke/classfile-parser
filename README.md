
# classfile-parser

This is Java classfile parser, written in pure Go.

## Usage

```go
import (
	"log"
	parser "github.com/wreulicke/classfile-parser"
)

func main() {
	f, err := os.Open("some/dir/foo.class")
	if err != nil {
		log.Fatal(err)
	}
	p := parser.New(f)
	classfile, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}
	// ...
}

```

## License

* MIT
