package yarn

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

const (
	yarn     = "yarn"
	lockfile = "yarn.lock"
)

type YarnAdapter struct{}

func (a YarnAdapter) FindLanguages(file *file.File) ([]*adapter.Language, error) {
	return nil, nil
}

func (a YarnAdapter) FindTools(file *file.File) ([]*adapter.Tool, error) {
	if isYarnFile(file) {
		tool := &adapter.Tool{
			Name:    yarn,
			Version: "",
		}

		return []*adapter.Tool{tool}, nil
	}

	return nil, nil
}

func (a YarnAdapter) FindDependencies(file *file.File) ([]*adapter.Dependency, error) {
	return nil, nil
}

func isYarnFile(file *file.File) bool {
	return file.Name() == lockfile
}
