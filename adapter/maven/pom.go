package maven

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/crqra/whatami/adapter"
)

type POM struct {
	XMLName              xml.Name             `xml:"project"`
	Version              string               `xml:"version"`
	Parent               *Parent              `xml:"parent"`
	Properties           *Properties          `xml:"properties"`
	DependencyManagement DependencyManagement `xml:"dependencyManagement"`
}

func readPOM(absPath string) (*POM, error) {
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	var pom POM
	if err := xml.Unmarshal(data, &pom); err != nil {
		panic(err)
	}

	return &pom, nil
}

func (pom POM) JavaVersion() string {
	for _, prop := range pom.Properties.Property {
		if prop.XMLName.Local == "java.version" {
			return string(prop.Value)
		}
	}

	return ""
}

func (pom POM) Dependencies() []*adapter.Dependency {
	deps := []*adapter.Dependency{}

	if pom.Parent != nil {
		deps = append(deps, &adapter.Dependency{
			Name:    fmt.Sprintf("%s.%s", pom.Parent.GroupID, pom.Parent.ArtifactID),
			Version: pom.Parent.Version,
			Type:    "parent",
		})
	}

	for _, dep := range pom.DependencyManagement.Dependencies.Dependencies {
		deps = append(deps, &adapter.Dependency{
			Name:    fmt.Sprintf("%s.%s", dep.GroupID, dep.ArtifactID),
			Version: dep.Version,
			Type:    dep.Scope,
		})
	}

	return deps
}

type Parent struct {
	XMLName    xml.Name `xml:"parent"`
	GroupID    string   `xml:"groupId"`
	ArtifactID string   `xml:"artifactId"`
	Version    string   `xml:"version"`
}

type Properties struct {
	XMLName  xml.Name   `xml:"properties"`
	Property []Property `xml:",any"`
}

type Property struct {
	XMLName xml.Name
	Value   []byte `xml:",chardata"`
}

type DependencyManagement struct {
	XMLName      xml.Name     `xml:"dependencyManagement"`
	Dependencies Dependencies `xml:"dependencies"`
}

type Dependencies struct {
	XMLName      xml.Name     `xml:"dependencies"`
	Dependencies []Dependency `xml:"dependency"`
}

type Dependency struct {
	XMLName    xml.Name `xml:"dependency"`
	GroupID    string   `xml:"groupId"`
	ArtifactID string   `xml:"artifactId"`
	Version    string   `xml:"version"`
	Scope      string   `xml:"scope"`
}
