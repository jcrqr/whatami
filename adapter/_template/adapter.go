package template

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

type TemplateAdapter struct{}

func (a TemplateAdapter) FindLanguages(*file.File) ([]*adapter.Language, error) {
	return nil, nil
}

func (a TemplateAdapter) FindTools(*file.File) ([]*adapter.Tool, error) {
	return nil, nil
}

func (a TemplateAdapter) FindDependencies(*file.File) ([]*adapter.Dependency, error) {
	return nil, nil
}
