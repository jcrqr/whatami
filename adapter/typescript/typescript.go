package typescript

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

const (
	ts       = "typescript"
	tsExt    = ".ts"
	tsxExt   = ".tsx"
	tsc      = "tsc"
	tsConfig = "tsconfig.json"
)

type TypeScriptAdapter struct{}

func (a TypeScriptAdapter) FindLanguages(file *file.File) ([]*adapter.Language, error) {
	if isTypeScriptFile(file) {
		lang := &adapter.Language{
			Name:    ts,
			Version: "",
		}

		return []*adapter.Language{lang}, nil
	}

	return nil, nil
}

func (a TypeScriptAdapter) FindTools(file *file.File) ([]*adapter.Tool, error) {
	if isTypeScriptFile(file) {
		tool := &adapter.Tool{
			Name:    tsc,
			Version: "",
		}

		return []*adapter.Tool{tool}, nil
	}

	return nil, nil
}

func (a TypeScriptAdapter) FindDependencies(file *file.File) ([]*adapter.Dependency, error) {
	return nil, nil
}

func isTypeScriptFile(file *file.File) bool {
	return file.Ext() == tsExt || file.Ext() == tsxExt || file.Name() == tsConfig
}
