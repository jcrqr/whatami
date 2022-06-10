package golang

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

const (
	goLang = "go"
	goMod  = "go.mod"
	goSum  = "go.sum"
	goExt  = ".go"
)

type GolangAdapter struct{}

func (a GolangAdapter) FindLanguages(f *file.File) ([]*adapter.Language, error) {
	lang := &adapter.Language{Name: goLang, Version: ""}

	if isGoMod(f) {
		mod, err := NewModFile(f.Path)
		if err != nil {
			return nil, err
		}

		lang.Version = mod.GoVersion()

		return []*adapter.Language{lang}, nil
	}

	if isGoFile(f) {
		return []*adapter.Language{lang}, nil
	}

	return nil, nil
}

func (a GolangAdapter) FindTools(f *file.File) ([]*adapter.Tool, error) {
	tool := &adapter.Tool{Name: goLang, Version: ""}

	if isGoMod(f) {
		mod, err := NewModFile(f.Path)
		if err != nil {
			return nil, err
		}

		tool.Version = mod.GoVersion()

		return []*adapter.Tool{tool}, nil
	}

	return nil, nil
}

func (a GolangAdapter) FindDependencies(f *file.File) ([]*adapter.Dependency, error) {
	if isGoMod(f) {
		mod, err := NewModFile(f.Path)
		if err != nil {
			return nil, err
		}

		return mod.Dependencies(), nil
	}

	return nil, nil
}

func isGoMod(f *file.File) bool {
	return f.Name() == goMod
}

func isGoFile(f *file.File) bool {
	return isGoMod(f) || f.Ext() == goExt
}
