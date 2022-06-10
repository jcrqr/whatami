package maven

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

const (
	mavenTool   = "maven"
	javaLang    = "java"
	pomFilename = "pom.xml"
)

type MavenAdapter struct{}

func (a MavenAdapter) FindLanguages(file *file.File) ([]*adapter.Language, error) {
	if file.Name() == pomFilename {
		pom, err := NewPOM(file.Path)
		if err != nil {
			return nil, err
		}

		lang := &adapter.Language{
			Name:    javaLang,
			Version: pom.JavaVersion(),
		}

		return []*adapter.Language{lang}, nil
	}

	return nil, nil
}

func (a MavenAdapter) FindTools(file *file.File) ([]*adapter.Tool, error) {
	if file.Name() == pomFilename {
		return []*adapter.Tool{{Name: mavenTool}}, nil
	}

	return nil, nil
}

func (a MavenAdapter) FindDependencies(file *file.File) ([]*adapter.Dependency, error) {
	if file.Name() != pomFilename {
		return nil, nil
	}

	pom, err := NewPOM(file.Path)
	if err != nil {
		return nil, err
	}

	return pom.Dependencies(), nil
}
