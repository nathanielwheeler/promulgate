package generate

import (
	"os"
	"path/filepath"
	"text/template"
)

// GetFuncmap will return all functions needed to parse template files into projects
func (p *Project) GetFuncmap() template.FuncMap {
	fm := template.FuncMap{
    "dir": p.Dir,
    "cmdnamedir": p.CMDNameDir,
    "license": p.License,
	}

	return fm
}

// Dir will create a directory given a string
func (p *Project) Dir(name string) {
	os.Mkdir(filepath.Join(p.Root, name), 0755)
}

// CMDNameDir will create a cmd/name directory for the project
func (p *Project) CMDNameDir() {
  os.MkdirAll(filepath.Join(p.Root, "cmd", p.Name), 0755)
}

// License will create a license file in the project root.
func (p *Project) License(name string) {
  switch name {
  case "mit":
    // TODO implement
    return
  default:
    return
  }
}