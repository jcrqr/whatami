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

func (a NPMAdapter) FindLanguages(f *file.File) ([]*adapter.Language, error) {
	return nil, nil
}

func (a NPMAdapter) FindTools(f *file.File) ([]*adapter.Tool, error) {
	if isNPMFile(f) {
		tools := []*adapter.Tool{
			{Name: nodeTool, Version: ""},
			{Name: npmTool, Version: ""},
		}

		return tools, nil
	}

	return nil, nil
}

func (a NPMAdapter) FindDependencies(f *file.File) ([]*adapter.Dependency, error) {
	if f.Name() != pkgFile {
		return nil, nil
	}

	pkg, err := NewPKG(f.Path)
	if err != nil {
		return nil, err
	}

	return pkg.Dependencies(), nil
}

func isNPMFile(f *file.File) bool {
	return f.Name() == pkgFile || f.Name() == lockfile
}
