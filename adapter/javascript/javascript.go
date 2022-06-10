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

func (a JavaScriptAdapter) FindLanguages(f *file.File) ([]*adapter.Language, error) {
	if f.Ext() == jsExt || f.Ext() == jsxExt {
		lang := &adapter.Language{
			Name:    js,
			Version: "",
		}

		return []*adapter.Language{lang}, nil
	}

	return nil, nil
}

func (a JavaScriptAdapter) FindTools(*file.File) ([]*adapter.Tool, error) {
	return nil, nil
}

func (a JavaScriptAdapter) FindDependencies(*file.File) ([]*adapter.Dependency, error) {
	return nil, nil
}
