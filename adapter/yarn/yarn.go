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

func (a YarnAdapter) FindLanguages(*file.File) ([]*adapter.Language, error) {
	return nil, nil
}

func (a YarnAdapter) FindTools(f *file.File) ([]*adapter.Tool, error) {
	if isYarnFile(f) {
		tool := &adapter.Tool{
			Name:    yarn,
			Version: "",
		}

		return []*adapter.Tool{tool}, nil
	}

	return nil, nil
}

func (a YarnAdapter) FindDependencies(*file.File) ([]*adapter.Dependency, error) {
	return nil, nil
}

func isYarnFile(f *file.File) bool {
	return f.Name() == lockfile
}
