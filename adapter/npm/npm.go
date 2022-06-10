package npm

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

const (
	npmTool  = "npm"
	nodeTool = "node"
	pkgFile  = "package.json"
	lockfile = "package-lock.json"
)

type NPMAdapter struct{}

func (a NPMAdapter) FindLanguages(file *file.File) ([]*adapter.Language, error) {
	return nil, nil
}

func (a NPMAdapter) FindTools(file *file.File) ([]*adapter.Tool, error) {
	if isNPMFile(file) {
		tools := []*adapter.Tool{
			{Name: nodeTool, Version: ""},
			{Name: npmTool, Version: ""},
		}

		return tools, nil
	}

	return nil, nil
}

func (a NPMAdapter) FindDependencies(file *file.File) ([]*adapter.Dependency, error) {
	if file.Name() != pkgFile {
		return nil, nil
	}

	pkg, err := NewPKG(file.Path)
	if err != nil {
		return nil, err
	}

	return pkg.Dependencies(), nil
}

func isNPMFile(file *file.File) bool {
	return file.Name() == pkgFile || file.Name() == lockfile
}
