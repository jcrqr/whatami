package maven

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

const (
	maven = "maven"
	pom   = "pom.xml"
)

type MavenAdapter struct{}

func (a MavenAdapter) FindLanguages(file *file.File) ([]*adapter.Language, error) {
	if file.Name() == pom {
		pom, err := readPOM(file.Path)
		if err != nil {
			return nil, err
		}

		lang := &adapter.Language{
			Name:    "java",
			Version: pom.JavaVersion(),
		}

		return []*adapter.Language{lang}, nil
	}

	return nil, nil
}

func (a MavenAdapter) FindTools(file *file.File) ([]*adapter.Tool, error) {
	if file.Name() == pom {
		return []*adapter.Tool{{Name: maven}}, nil
	}

	return nil, nil
}

func (a MavenAdapter) FindDependencies(file *file.File) ([]*adapter.Dependency, error) {
	if file.Name() != pom {
		return nil, nil
	}

	pom, err := readPOM(file.Path)
	if err != nil {
		return nil, err
	}

	return pom.Dependencies(), nil
}
