package npm

import (
	"os"

	"github.com/antchfx/jsonquery"
	"github.com/crqra/whatami/adapter"
)

type PKG struct {
	doc *jsonquery.Node
}

func NewPKG(path string) (*PKG, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	doc, err := jsonquery.Parse(f)
	if err != nil {
		return nil, err
	}

	return &PKG{doc: doc}, nil
}

func (pkg PKG) Dependencies() []*adapter.Dependency {
	var (
		deps          = []*adapter.Dependency{}
		rawDeps, _    = jsonquery.QueryAll(pkg.doc, "/dependencies/*")
		rawDevDeps, _ = jsonquery.QueryAll(pkg.doc, "/devDependencies/*")
	)

	for _, dep := range rawDeps {
		deps = append(deps, &adapter.Dependency{
			Name:    dep.Data,
			Version: dep.InnerText(),
			Type:    "production",
		})
	}

	for _, dep := range rawDevDeps {
		deps = append(deps, &adapter.Dependency{
			Name:    dep.Data,
			Version: dep.InnerText(),
			Type:    "development",
		})
	}

	return deps
}
