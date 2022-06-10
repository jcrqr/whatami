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
	FindTools(*file.File) ([]*Tool, error)
	FindDependencies(*file.File) ([]*Dependency, error)
	FindLanguages(*file.File) ([]*Language, error)
}
