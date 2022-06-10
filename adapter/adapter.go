package adapter

import "github.com/crqra/whatami/file"

type Language struct {
	Name    string `json:"-"`
	Version string `json:"version"`
}

type Tool struct {
	Name    string `json:"-"`
	Version string `json:"version"`
}

type Dependency struct {
	Name    string `json:"-"`
	Version string `json:"version"`
	Type    string `json:"type"`
}

type Adapter interface {
	FindTools(*file.File) ([]*Tool, error)
	FindDependencies(*file.File) ([]*Dependency, error)
	FindLanguages(*file.File) ([]*Language, error)
}
