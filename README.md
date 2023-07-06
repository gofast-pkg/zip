# zip

[![Static Badge](https://img.shields.io/badge/project%20use%20codesystem-green?link=https%3A%2F%2Fgithub.com%2Fgofast-pkg%2Fcodesystem)](https://github.com/gofast-pkg/codesystem)
![Build](https://github.com/gofast-pkg/zip/actions/workflows/ci.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/gofast-pkg/zip.svg)](https://pkg.go.dev/github.com/gofast-pkg/zip)
[![codecov](https://codecov.io/gh/gofast-pkg/zip/branch/main/graph/badge.svg?token=7TCE3QB21E)](https://codecov.io/gh/gofast-pkg/zip)
[![Release](https://img.shields.io/github/release/gofast-pkg/zip?style=flat-square)](https://github.com/gofast-pkg/zip/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofast-pkg/zip)](https://goreportcard.com/report/github.com/gofast-pkg/zip)
[![codebeat badge](https://codebeat.co/badges/9338570f-6fe5-4095-bf2f-93a53c5dc800)](https://codebeat.co/projects/github-com-gofast-pkg-zip-main)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/gofast-pkg/zip/blob/main/LICENSE)

Package zip process a compressed file read, write and clone it

## Install

``` bash
$> go get github.com/gofast-pkg/zip@latest
```

## Usage

``` Golang
import github.com/gofast-pkg/zip

func main() {
  file, err := os.Open(testReadCloserZipFile)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  r, err := zip.NewReader(file)
  if err != nil {
    log.Fatal(err)
  }

  // Iterate through the files in the archive,
  for i := 0; i < r.NFiles(); i++ {
    var content []byte
    if content, err = r.ContentFile(i); err != nil {
      panic(err)
    }
    // do something
  }
}
```

Check the [go documentation](https://pkg.go.dev/github.com/gofast-pkg/zip) for more details.

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
