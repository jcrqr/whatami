package golang

import (
	"os"

	"github.com/crqra/whatami/adapter"
	"golang.org/x/mod/modfile"
)

type ModFile struct {
	mod *modfile.File
}

func NewModFile(path string) (*ModFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	mod, err := modfile.Parse(path, data, nil)
	if err != nil {
		return nil, err
	}

	return &ModFile{mod: mod}, nil
}

func (mod ModFile) GoVersion() string {
	return mod.mod.Go.Version
}

func (mod ModFile) Dependencies() []*adapter.Dependency {
	deps := []*adapter.Dependency{}

	for _, req := range mod.mod.Require {
		t := "direct"

		if req.Indirect {
			t = "indirect"
		}

		deps = append(deps, &adapter.Dependency{
			Name:    req.Mod.Path,
			Version: req.Mod.Version,
			Type:    t,
		})
	}

	return deps
}
