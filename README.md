# cmdr-go-starter

[![Go](https://github.com/hedzr/cmdr-go-starter/actions/workflows/go.yml/badge.svg)](https://github.com/hedzr/cmdr-go-starter/actions/workflows/go.yml)

A template repository to build your first Golang Open Source app based [cmdr](https://github.com/hedzr/cmdr).

> powered by [cmdr](https://github.com/hedzr/cmdr) v1.11+.

## fast guide:

1. All you have to do is click the <kbd>Use this template</kbd> button upon this page.
2. run! (`go run ./cli/app/cli/app/main.go`)

## with command-line:

### new repo

1. clone cmdr-go-starter as a template
   ```bash
   # clone cmdr-go-starter as a template
   git clone https://github.com/hedzr/cmdr-go-starter.git new-repo
   cd new-repo
   git push git@github.com:yourname/new-repo.git +master:master
   ```

2. clone the `new-repo` to your working directory:
   ```bash
   # in the working directory of your new-repo
   cd ~/work
   git clone git@github.com:yourname/new-repo.git
   cd new-repo
   ```

3. do rename stuffs ...



## Getting Started (For the generated golang project)

To run the CLI app:

```bash
# go install -v github.com/swaggo/swag/cmd/swag
go generate ./...               # run it once at least, for gen the swagger-doc files from skeletons
go run ./cli/your-app/main.go   # build the CLI app
```

## Status

- cmdr v1.11.26+
- 

## LICENSE

Apache 2.0
