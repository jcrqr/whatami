package main

import (
	"encoding/json"
	"flag"
	"io/fs"
	"log"
	"os"
	"regexp"

	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

type output struct {
	Tools        map[string]*adapter.Tool       `json:"tools"`
	Dependencies map[string]*adapter.Dependency `json:"dependencies"`
	Languages    map[string]*adapter.Language   `json:"languages"`
}

type options struct {
	showUsage bool
	root      string
	ignored   filterFlag
	adapters  []adapter.Adapter
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	opts := options{
		root:     cwd,
		adapters: registeredAdapters(),
	}

	fs := flag.NewFlagSet("whatami", flag.ExitOnError)

	fs.BoolVar(&opts.showUsage, "h", false, "Show usage")
	fs.Var(&opts.ignored, "i", "Patterns to ignore")

	fs.Parse(os.Args[1:])

	if opts.showUsage {
		fs.Usage()
		os.Exit(0)
	}

	if path := fs.Arg(0); path != "" {
		opts.root = path
	}

	logger := log.New(os.Stderr, "whatami: ", 0)

	out, err := execute(opts)
	if err != nil {
		logger.Fatalf("Error: %s\n", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	if err := enc.Encode(out); err != nil {
		panic(err)
	}
}

func execute(opts options) (*output, error) {
	out := &output{
		Tools:        map[string]*adapter.Tool{},
		Dependencies: map[string]*adapter.Dependency{},
		Languages:    map[string]*adapter.Language{},
	}

	ignoredRe := []*regexp.Regexp{}

	for _, expr := range opts.ignored {
		re, err := regexp.Compile(expr)
		if err != nil {
			return nil, err
		}

		ignoredRe = append(ignoredRe, re)
	}

	err := file.Walk(opts.root, func(file *file.File) error {
		for _, re := range ignoredRe {
			if re.MatchString(file.Path) {
				return fs.SkipDir
			}
		}

		for _, a := range opts.adapters {
			if err := findAndUpdate(a, file, out); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return out, nil
}

func findAndUpdate(adapter adapter.Adapter, file *file.File, out *output) error {
	tools, err := adapter.FindTools(file)
	if err != nil {
		return err
	}

	for _, tool := range tools {
		// If tool is already defined and has a version, skip
		if t, ok := out.Tools[tool.Name]; ok && t.Version != "" {
			continue
		}

		out.Tools[tool.Name] = tool
	}

	deps, err := adapter.FindDependencies(file)
	if err != nil {
		return err
	}

	for _, dep := range deps {
		// If dependency is already defined and has a version, skip
		if existingDep, ok := out.Languages[dep.Name]; ok && existingDep.Version != "" {
			continue
		}

		out.Dependencies[dep.Name] = dep
	}

	langs, err := adapter.FindLanguages(file)
	if err != nil {
		return err
	}

	for _, lang := range langs {
		// If language is already defined and has a version, skip
		if existingLang, ok := out.Languages[lang.Name]; ok && existingLang.Version != "" {
			continue
		}

		out.Languages[lang.Name] = lang
	}

	return nil
}
