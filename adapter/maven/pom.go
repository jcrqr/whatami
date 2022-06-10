package maven

import (
	"fmt"
	"os"
	"regexp"

	"github.com/antchfx/xmlquery"
	"github.com/crqra/whatami/adapter"
)

type POM struct {
	doc *xmlquery.Node
}

func NewPOM(path string) (*POM, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	doc, err := xmlquery.Parse(f)
	if err != nil {
		return nil, err
	}

	return &POM{doc: doc}, nil
}

func (pom POM) JavaVersion() string {
	prop := xmlquery.FindOne(pom.doc, "/project/properties/java.version")
	if prop != nil {
		return prop.InnerText()
	}

	return ""
}

func (pom POM) Dependencies() []*adapter.Dependency {
	var (
		deps    = []*adapter.Dependency{}
		rawDeps = xmlquery.Find(pom.doc, "/project/dependencyManagement/dependencies/dependency")
	)

	for _, dep := range rawDeps {
		deps = append(deps, parseDependencyElement(pom.doc, dep))
	}

	rawParent := xmlquery.FindOne(pom.doc, "/project/parent")
	if rawParent != nil {
		dep := parseDependencyElement(pom.doc, rawParent)
		dep.Type = "parent"

		deps = append(deps, dep)
	}

	return deps
}

func parseDependencyElement(doc *xmlquery.Node, dep *xmlquery.Node) *adapter.Dependency {
	var (
		groupId    string
		artifactId string
		version    string
		scope      string
	)

	if id := dep.SelectElement("groupId"); id != nil {
		groupId = id.InnerText()
	}

	if id := dep.SelectElement("artifactId"); id != nil {
		artifactId = id.InnerText()
	}

	if v := dep.SelectElement("version"); v != nil {
		version = findVersion(doc, v.InnerText())
	}

	if s := dep.SelectElement("scope"); s != nil {
		scope = s.InnerText()
	}

	return &adapter.Dependency{
		Name:    fmt.Sprintf("%s.%s", groupId, artifactId),
		Version: version,
		Type:    scope,
	}
}

func findVersion(doc *xmlquery.Node, version string) string {
	re := regexp.MustCompile("\\${(.*)}")
	matches := re.FindStringSubmatch(version)

	//nolint:gomnd
	if len(matches) != 2 {
		return version
	}

	vProp := xmlquery.FindOne(doc, "/project/properties/"+matches[1])
	if vProp == nil {
		return version
	}

	return vProp.InnerText()
}
