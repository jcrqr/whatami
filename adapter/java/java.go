package java

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

const (
	langName = "java"
	validExt = ".java"
)

type JavaAdapter struct{}

func (a JavaAdapter) FindLanguages(file *file.File) ([]*adapter.Language, error) {
	if file.Ext() == validExt {
		return []*adapter.Language{{Name: langName}}, nil
	}

	return nil, nil
}

func (a JavaAdapter) FindTools(*file.File) ([]*adapter.Tool, error) {
	return nil, nil
}

func (a JavaAdapter) FindDependencies(*file.File) ([]*adapter.Dependency, error) {
	return nil, nil
}
