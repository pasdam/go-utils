# go-utils

[![Go Report Card](https://goreportcard.com/badge/github.com/pasdam/go-utils)](https://goreportcard.com/report/github.com/pasdam/go-utils)
[![CI Status](https://github.com/pasdam/go-utils/workflows/Continuous%20integration/badge.svg)](https://github.com/pasdam/go-utils/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/github.com/pasdam/go-utils?status.svg)](https://godoc.org/github.com/pasdam/go-utils)

A collection of Go utility packages for common development tasks, primarily
focused on testing utilities and data validation.

## Packages

### [assertutils](pkg/assertutils)

Provides testing utilities for comparing errors in test assertions, making
error comparison more reliable and consistent in unit tests.

### [filetestutils](pkg/filetestutils)

Contains helpers for file system operations in tests, including creating
temporary files and directories, and utilities for checking file existence
and content.

### [jsonschemautils](pkg/jsonschemautils)

Utilities for validating data against JSON schemas, providing a simple
interface for schema validation of structured data.

### [yamlutils](pkg/yamlutils)

Tools for converting YAML data to JSON, useful when working with
configuration files and data interchange formats.

## Usage

To use these utilities in your project:

```bash
go get github.com/pasdam/go-utils
```

Then import the specific package you need:

```go
import "github.com/pasdam/go-utils/pkg/filetestutils"
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE)
file for details.
