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

func (a MavenAdapter) FindLanguages(f *file.File) ([]*adapter.Language, error) {
	if f.Name() == pomFilename {
		pom, err := NewPOM(f.Path)
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

func (a MavenAdapter) FindTools(f *file.File) ([]*adapter.Tool, error) {
	if f.Name() == pomFilename {
		return []*adapter.Tool{{Name: mavenTool}}, nil
	}

	return nil, nil
}

func (a MavenAdapter) FindDependencies(f *file.File) ([]*adapter.Dependency, error) {
	if f.Name() != pomFilename {
		return nil, nil
	}

	pom, err := NewPOM(f.Path)
	if err != nil {
		return nil, err
	}

	return pom.Dependencies(), nil
}
