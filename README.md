# zip

[![Static Badge](https://img.shields.io/badge/project%20use%20codesystem-green?link=https%3A%2F%2Fgithub.com%2Fgofast-pkg%2Fcodesystem)](https://github.com/gofast-pkg/codesystem)
![Build](https://github.com/gofast-pkg/zip/actions/workflows/ci.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/gofast-pkg/zip.svg)](https://pkg.go.dev/github.com/gofast-pkg/zip)
[![codecov](https://codecov.io/gh/gofast-pkg/zip/branch/main/graph/badge.svg?token=7TCE3QB21E)](https://codecov.io/gh/gofast-pkg/zip)
[![Release](https://img.shields.io/github/release/gofast-pkg/zip?style=flat-square)](https://github.com/gofast-pkg/zip/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofast-pkg/zip)](https://goreportcard.com/report/github.com/gofast-pkg/zip)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/gofast-pkg/zip/blob/main/LICENSE)

This package provides a simple abstraction to read and process files from a ZIP archive.

It exposes a Reader interface that allows consumers to interact with compressed files without dealing directly with the underlying archive/zip implementation. The package is designed to be straightforward, testable, and easy to integrate into your codebase.

## Install

``` bash
$> go get github.com/gofast-pkg/zip@latest
```

## Key Features

* Unified interface (Reader)
* Retrieve the number of files in the archive
* Read file contents by index
* Access file metadata
* Write file contents directly to an io.Writer
* In-memory processing

Files are accessed by index, ensuring deterministic behavior aligned with the ZIP structure.

## Usage

Check [the example](./example_test.go) and the [go documentation](https://pkg.go.dev/github.com/gofast-pkg/zip) for more details.

## Contributing

&nbsp;:grey_exclamation:&nbsp; Use issues for everything

Read more informations with the [CONTRIBUTING_GUIDE](./.github/CONTRIBUTING.md)

For all changes, please update the CHANGELOG.txt file by replacing the existant content.

Thank you &nbsp;:pray:&nbsp;&nbsp;:+1:&nbsp;

<a href="https://github.com/gofast-pkg/zip/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=gofast-pkg/zip" />
</a>

Made with [contrib.rocks](https://contrib.rocks).

## Licence

[MIT](https://github.com/gofast-pkg/zip/blob/main/LICENSE)
