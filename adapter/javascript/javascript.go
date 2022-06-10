package javascript

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

const (
	js     = "javascript"
	jsExt  = ".js"
	jsxExt = ".jsx"
)

type JavaScriptAdapter struct{}

func (a JavaScriptAdapter) FindLanguages(file *file.File) ([]*adapter.Language, error) {
	if file.Ext() == jsExt || file.Ext() == jsxExt {
		lang := &adapter.Language{
			Name:    js,
			Version: "",
		}

		return []*adapter.Language{lang}, nil
	}

	return nil, nil
}

func (a JavaScriptAdapter) FindTools(file *file.File) ([]*adapter.Tool, error) {
	return nil, nil
}

func (a JavaScriptAdapter) FindDependencies(file *file.File) ([]*adapter.Dependency, error) {
	return nil, nil
}
