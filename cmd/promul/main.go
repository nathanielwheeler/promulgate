package main

import (
	"fmt"
  "os"
  
  "promulgate/pkg/generate"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(1)
	}
}

func run() error {
  p := generate.Project{
    Name: "test",
    Template: "default",
  }

  err := p.Generate()
  return err
}
