package terraform

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

const tfExt = ".tf"

var (
	// nolint: gochecknoglobals
	tfTool = adapter.Tool{
		Name:    "terraform",
		Version: "",
	}

	// nolint: gochecknoglobals
	hclLang = adapter.Language{
		Name:    "hcl",
		Version: "",
	}
)

type TerraformAdapter struct{}

func (a TerraformAdapter) FindLanguages(f *file.File) ([]*adapter.Language, error) {
	if isTerraformFile(f) {
		return []*adapter.Language{&hclLang}, nil
	}

	return nil, nil
}

func (a TerraformAdapter) FindTools(f *file.File) ([]*adapter.Tool, error) {
	if isTerraformFile(f) {
		return []*adapter.Tool{&tfTool}, nil
	}

	return nil, nil
}

func (a TerraformAdapter) FindDependencies(*file.File) ([]*adapter.Dependency, error) {
	return nil, nil
}

func isTerraformFile(f *file.File) bool {
	return f.Ext() == tfExt
}
