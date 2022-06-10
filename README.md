# whatami

> What Am I? - point to a project and see what tools, languages and dependencies
> it has.

## About

### What?

Tool that recursively walks a directory and based on the files present derives:

- Tools: package managers, compilers, linters, etc.
- Languages: markup, configuration and programming languages
- Dependencies: software dependencies from package managers configuration files

For each of the above a version _may_ be available in which case it's included.

### Why?

My main use case is to use the information outputted to create dynamic CI/CD
pipelines. Based on the output, I know what tools I need to have in place for a
particular project and can even infer what commands I can run to do certain tasks
like building or testing the project.

### How?

For each file or directory found while walking, a set of pre-defined _adapters_
is executed to derive tools, languages and dependencies.

If more than one adapter finds the same tools, languages or dependencies, they're
de-duplicated based on the version. Wins the first tool, language or dependency
that for which a version was found (regardless if subsequent findings also have
a version).

## Installing

```console
$ go install github.com/crqra/whatami/cmd/whatami@latest
```

## Usage

```
$ whatami -h
usage: whatami [-h] [-i=<PATH> ...] [directory]
```

### Flags

| Flag        | Description                           |
| ----------- | ------------------------------------- |
| `-h`        | Show usage                            |
| `-i <PATH>` | Path patterns to ignore. Accepts many |

## Example

In a standard TypeScript project:

```bash
$ ls
index.ts node_modules/ package-lock.json package.json

# Run whatami
$ whatami -i node_modules
```

<details>
<summary>See output</summary>

```json
{
  "tools": {
    "node": {},
    "npm": {},
    "tsc": {
      "version": "^4.7.3"
    }
  },
  "dependencies": {
    "express": {
      "version": "^4.18.1",
      "type": "production"
    },
    "typescript": {
      "version": "^4.7.3",
      "type": "development"
    }
  },
  "languages": {
    "typescript": {}
  }
}
```

</details>

## License

This project is released under the [MIT License](LICENSE).
