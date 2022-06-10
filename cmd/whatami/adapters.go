package main

import (
	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/adapter/docker"
	"github.com/crqra/whatami/adapter/golang"
	"github.com/crqra/whatami/adapter/hcl"
	"github.com/crqra/whatami/adapter/java"
	"github.com/crqra/whatami/adapter/javascript"
	"github.com/crqra/whatami/adapter/maven"
	"github.com/crqra/whatami/adapter/npm"
	"github.com/crqra/whatami/adapter/taskfile"
	"github.com/crqra/whatami/adapter/terraform"
	"github.com/crqra/whatami/adapter/typescript"
	"github.com/crqra/whatami/adapter/yarn"
)

func registeredAdapters() []adapter.Adapter {
	return []adapter.Adapter{
		&docker.DockerAdapter{},
		&golang.GolangAdapter{},
		&hcl.HCLAdapter{},
		&java.JavaAdapter{},
		&javascript.JavaScriptAdapter{},
		&maven.MavenAdapter{},
		&npm.NPMAdapter{},
		&taskfile.TaskfileAdapter{},
		&terraform.TerraformAdapter{},
		&typescript.TypeScriptAdapter{},
		&yarn.YarnAdapter{},
	}
}
