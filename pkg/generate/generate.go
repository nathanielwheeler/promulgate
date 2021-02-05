package generate

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Project holds the metadata needed to generate a new project.
type Project struct {
	Name     string
	Template string
}

// Generate will take a project and generate it.
func (p *Project) Generate() error {
	// Validate root path
	wd, err := os.Getwd()
	if err != nil {
		return errors.New("couldn't get working directory")
	}
	root := filepath.Join(wd, strings.ToLower(p.Name))

	// Validate template
	_, b, _, _ := runtime.Caller(0)
  tplFile := filepath.Join(filepath.Dir(b), "..", "..", "templates", p.Template)
  tplFile = tplFile + ".tpl"
  _, err = os.Stat(tplFile)
  if err != nil {
    return err
  }

  // Generate project root directory
  err = os.Mkdir(root, 0755)
  if err != nil {
    return err
  }

  // TODO generate template functions and go through template

	return nil
}
