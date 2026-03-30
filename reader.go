// Package zip process a compressed file / body.
package zip

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
)

// Errors supported by the zip.Reader operations.
var (
	ErrNilReader         = errors.New("reader is nil")
	ErrNilWriter         = errors.New("writer is nil")
	ErrToReadReader      = errors.New("fails to read the reader")
	ErrToCreateZipReader = errors.New("fails to create a zip.Reader")
	ErrInvalidIndex      = errors.New("file index reference is invalid")
	ErrToReadFile        = errors.New("fails to read file in the zip.Reader")
	ErrToWriteFile       = errors.New("fails to write file")
	ErrToOpenFile        = errors.New("fails to open file in the zip.Reader")
)

// Reader is the interface that wraps the basic methods to process compressed file(s).
//
//mockery:generate: true
type Reader interface {
	// NumFile returns the number of files in the compressed file / body.
	NumFile() int
	// Read the content of a file and return the body.
	Read(index int) ([]byte, error)
	// InfoFile returns the information of a file.
	InfoFile(index int) (fs.FileInfo, error)
	// Write the content in the io.Writer parameter.
	Write(w io.Writer, index int) error
}

type reader struct {
	input *zip.Reader
	copy  []byte
}

// NewReader returns a Reader interface for the zip.
// It receives a io.ReadCloser interface from http request or file.
func NewReader(input io.ReadCloser) (Reader, error) {
	var err error
	var data []byte
	var zr *zip.Reader

	if input == nil {
		return nil, ErrNilReader
	}
	if data, err = io.ReadAll(input); err != nil {
		return nil, errors.Join(err, ErrToReadReader)
	}

	if zr, err = zip.NewReader(bytes.NewReader(data), int64(len(data))); err != nil {
		return nil, errors.Join(err, ErrToCreateZipReader)
	}

	return &reader{
		input: zr,
		copy:  data,
	}, nil
}

func (r reader) InfoFile(index int) (fs.FileInfo, error) {
	var err error
	var info fs.FileInfo

	if err = r.indexValidation(index); err != nil {
		return nil, err
	}
	info = r.input.File[index].FileInfo()

	return info, nil
}

func (r reader) NumFile() int {
	return len(r.input.File)
}

func (r reader) Read(index int) ([]byte, error) {
	var err error
	var rc io.ReadCloser

	if err = r.indexValidation(index); err != nil {
		return nil, err
	}

	if rc, err = r.input.File[index].Open(); err != nil {
		return nil, errors.Join(ErrToOpenFile, fmt.Errorf("index file %d: %w", index, err))
	}
	defer func() { err = errors.Join(err, rc.Close()) }()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.Join(ErrToReadFile, fmt.Errorf("index file %d: %w", index, err))
	}

	return body, nil
}

func (r reader) Write(w io.Writer, index int) error {
	var err error
	var body []byte

	if w == nil {
		return ErrNilWriter
	}

	if err = r.indexValidation(index); err != nil {
		return err
	}
	if body, err = r.Read(index); err != nil {
		return errors.Join(ErrToReadFile, fmt.Errorf("index file %d: %w", index, err))
	}
	if _, err := w.Write(body); err != nil {
		return errors.Join(ErrToWriteFile, fmt.Errorf("index file %d: %w", index, err))
	}

	return nil
}

func (r reader) indexValidation(index int) error {
	if index < 0 || index >= len(r.input.File) {
		return ErrInvalidIndex
	}

	return nil
}
