package adapter

import "github.com/crqra/whatami/file"

type Language struct {
	Name    string `json:"-"`
	Version string `json:"version,omitempty"`
}

type Tool struct {
	Name    string `json:"-"`
	Version string `json:"version,omitempty"`
}

type Dependency struct {
	Name    string `json:"-"`
	Version string `json:"version,omitempty"`
	Type    string `json:"type,omitempty"`
}

type Adapter interface {
	FindTools(f *file.File) ([]*Tool, error)
	FindDependencies(f *file.File) ([]*Dependency, error)
	FindLanguages(f *file.File) ([]*Language, error)
}
