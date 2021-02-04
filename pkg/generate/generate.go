package generate

// Project holds the metadata needed to generate a new project.  Using the NewProject function to construct a Project is recommended, as it will validate values automatically.
type Project struct {
	Name        string
	Destination string
	Template    string
}

// NewProject is a constructor function for Project that will validate all of that project's data.
func NewProject(name, destination, template string) (*Project, error) {

  p := Project{
    Name: name,
    Destination: destination,
    Template: template,
  }
  return &p, nil
}

// Generate will take a project and generate it.
func (p *Project) Generate() error {
  
}