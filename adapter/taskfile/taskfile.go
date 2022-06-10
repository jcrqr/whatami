package taskfile

import (
	"os"

	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
	"gopkg.in/yaml.v3"
)

const (
	taskTool = "task"
	taskfile = "Taskfile"
	ymlExt   = ".yml"
	yamlExt  = ".yaml"
)

type TaskfileAdapter struct{}

func (a TaskfileAdapter) FindLanguages(*file.File) ([]*adapter.Language, error) {
	return nil, nil
}

func (a TaskfileAdapter) FindTools(f *file.File) ([]*adapter.Tool, error) {
	if !isTaskfile(f) {
		return nil, nil
	}

	version, err := getVersion(f)
	if err != nil {
		return nil, err
	}

	tool := &adapter.Tool{
		Name:    taskTool,
		Version: version,
	}

	return []*adapter.Tool{tool}, nil
}

func (a TaskfileAdapter) FindDependencies(*file.File) ([]*adapter.Dependency, error) {
	return nil, nil
}

func isTaskfile(f *file.File) bool {
	return f.Name() == taskfile+ymlExt || f.Name() == taskfile+yamlExt
}

func getVersion(f *file.File) (string, error) {
	data, err := os.ReadFile(f.Path)
	if err != nil {
		return "", err
	}

	var file struct {
		Version string `yaml:"version"`
	}

	if err := yaml.Unmarshal(data, &file); err != nil {
		return "", err
	}

	return file.Version, nil
}
