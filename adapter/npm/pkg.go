package npm

import (
	"encoding/json"
	"os"

	"github.com/crqra/whatami/adapter"
)

type pkg struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func readPkg(path string) (*pkg, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var pkg *pkg

	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	return pkg, nil
}

func (pkg pkg) allDependencies() []*adapter.Dependency {
	deps := []*adapter.Dependency{}

	for name, version := range pkg.Dependencies {
		deps = append(deps, &adapter.Dependency{
			Name:    name,
			Version: version,
			Type:    "production",
		})
	}

	for name, version := range pkg.DevDependencies {
		deps = append(deps, &adapter.Dependency{
			Name:    name,
			Version: version,
			Type:    "development",
		})
	}

	return deps
}
