// Package zip process a compressed file / body
package zip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
)

// Reader is the interface that wraps the basic methods to process a compressed file / body
//
//go:generate mockery --name=Reader --output=mocks --filename=reader.go --outpkg=mocks
type Reader interface {
	Create(filepath string) error
	NFiles() int
	ContentFile(index int) ([]byte, error)
	InfoFile(index int) (fs.FileInfo, error)
	WriteFile(index int, filepath string) error
}

type reader struct {
	input *zip.Reader
	copy  []byte
}

// NewReader returns a Reader interface
// It receives a io.ReadCloser interface from http request or file
func NewReader(input io.ReadCloser) (Reader, error) {
	var err error

	if input == nil {
		return nil, ErrInvalidInput
	}
	r := reader{}
	if r.copy, err = io.ReadAll(input); err != nil {
		return nil, err
	}
	body := make([]byte, len(r.copy))
	copy(body, r.copy)

	if r.input, err = zip.NewReader(bytes.NewReader(body), int64(len(body))); err != nil {
		return nil, err
	}

	return &r, nil
}

// WriteFile writes the content of a file in a specific path
func (r reader) WriteFile(index int, filepath string) error {
	var err error
	var info fs.FileInfo
	var body []byte
	var f *os.File

	if err = r.indexValidation(index); err != nil {
		return err
	}
	if info, err = r.InfoFile(index); err != nil {
		return err
	}
	if body, err = r.ContentFile(index); err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s", filepath, info.Name())
	if f, err = os.Create(path); err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write(body); err != nil {
		return err
	}

	return nil
}

// InfoFile returns the information of a file
func (r reader) InfoFile(index int) (fs.FileInfo, error) {
	var err error
	var info fs.FileInfo

	if err = r.indexValidation(index); err != nil {
		return nil, err
	}
	info = r.input.File[index].FileInfo()

	return info, nil
}

// ContentFile returns the content of a file
func (r reader) ContentFile(index int) ([]byte, error) {
	var err error
	var rc io.ReadCloser

	if err = r.indexValidation(index); err != nil {
		return nil, err
	}

	if rc, err = r.input.File[index].Open(); err != nil {
		return nil, err
	}
	defer rc.Close()

	return io.ReadAll(rc)
}

// NFiles returns the number of files in the compressed file / body
func (r reader) NFiles() int {
	return len(r.input.File)
}

// Create a file in a specific path
func (r reader) Create(filepath string) error {
	var err error
	var f *os.File

	if f, err = os.Create(filepath); err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(r.copy); err != nil {
		return err
	}

	return nil
}

func (r reader) indexValidation(index int) error {
	if index < 0 || index >= len(r.input.File) {
		return ErrInvalidIndex
	}

	return nil
}
