# whatami

> What Am I? - point to a project and see what tools, languages and dependencies
> it has.

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

# License

This project is released under the [MIT License](LICENSE).
