package typescript

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/adapter/npm"
	"github.com/crqra/whatami/file"
)

const (
	tsLang      = "typescript"
	tsExt       = ".ts"
	tsxExt      = ".tsx"
	tscTool     = "tsc"
	tsConfig    = "tsconfig.json"
	pkgFilename = "package.json"
)

type TypeScriptAdapter struct{}

func (a TypeScriptAdapter) FindLanguages(f *file.File) ([]*adapter.Language, error) {
	lang := &adapter.Language{
		Name:    tsLang,
		Version: "",
	}

	if isTypeScriptFile(f) {
		return []*adapter.Language{lang}, nil
	}

	if isPackageFile(f) {
		dep, err := findDependency(f)
		if err != nil {
			return nil, err
		}

		if dep == nil {
			return nil, nil
		}

		lang.Version = dep.Version

		return []*adapter.Language{lang}, nil
	}

	return nil, nil
}

func (a TypeScriptAdapter) FindTools(f *file.File) ([]*adapter.Tool, error) {
	tool := &adapter.Tool{
		Name:    tscTool,
		Version: "",
	}

	if isPackageFile(f) {
		dep, err := findDependency(f)
		if err != nil {
			return nil, err
		}

		if dep != nil {
			tool.Version = dep.Version

			return []*adapter.Tool{tool}, nil
		}

		return nil, nil
	}

	if isTypeScriptFile(f) {
		return []*adapter.Tool{tool}, nil
	}

	return nil, nil
}

func (a TypeScriptAdapter) FindDependencies(*file.File) ([]*adapter.Dependency, error) {
	return nil, nil
}

func isTypeScriptFile(f *file.File) bool {
	return f.Ext() == tsExt || f.Ext() == tsxExt || f.Name() == tsConfig
}

func isPackageFile(f *file.File) bool {
	return f.Name() == pkgFilename
}

func findDependency(f *file.File) (*adapter.Dependency, error) {
	pkg, err := npm.NewPKG(f.Path)
	if err != nil {
		return nil, err
	}

	for _, dep := range pkg.Dependencies() {
		if dep.Name == tsLang {
			return dep, nil
		}
	}

	return nil, nil
}
