package generate

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

// Project holds the metadata needed to generate a new project.
type Project struct {
	Name     string
	Template string
	Root     string
}

// Generate will take a project and generate it.
func (p Project) Generate() error {
	// Validate root path
	wd, err := os.Getwd()
	if err != nil {
		return errors.New("couldn't get working directory")
	}
	p.Root = filepath.Join(wd, strings.ToLower(p.Name))

	// Validate template
	_, b, _, _ := runtime.Caller(0)
	tplFile := filepath.Join(filepath.Dir(b), "..", "..", "templates", p.Template)
	tplFile = tplFile + ".tpl"
	_, err = os.Stat(tplFile)
	if err != nil {
		return err
	}

	// Generate project root directory
	err = os.Mkdir(p.Root, 0755)
	if err != nil {
		return err
	}

	fm := p.GetFuncmap()

	tpl, err := template.New("").Funcs(fm).Parse(tplFile)
	if err != nil {
		return err
	}

	err = tpl.Execute(os.Stdout, p)

	return nil
}
