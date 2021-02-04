package promul

import (
	"fmt"
  "os"
  
  "github.com/nathanielwheeler/promulgator/pkg/generate"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(1)
	}
}

func run() error {
	generate.
}
