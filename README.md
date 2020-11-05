# Cliq CLI
---
`Quickstart for your CLI`

## Benefits
---
- Uses cobra and viper libraries.
- Support for color and no color modes.
- Sample config parameters with config file.

## Development
- Make sure you clone this repository in path `$GOPATH/src/github.com/go143/`
- Add your command in folder `cmd/`
- Add nested commands to your command easily. See `cmd/details.go`
- Add custom flags per command level. See `cmd/root.go`

## Installation
- `$ go install` From repository root.
- `$ cliq --help` :bomb:
