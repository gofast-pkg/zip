package zip

import (
	"archive/zip"
	"bytes"
	"os"
	"testing"

	"github.com/gofast-pkg/zip/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type failingWriter struct{}

func (f failingWriter) Write(_ []byte) (int, error) {
	return 0, assert.AnError
}

func TestNewReader(t *testing.T) {
	t.Run("Should return an error because reader is nil", func(t *testing.T) {
		r, err := NewReader(nil)
		require.ErrorIs(t, err, ErrNilReader)
		assert.Nil(t, r)
	})
	t.Run("Should return an error because io.ReadCloser is already closed", func(t *testing.T) {
		var err error
		var file *os.File

		file, err = os.Open(testdata.TestZipFile)
		require.NoError(t, err)

		err = file.Close()
		require.NoError(t, err)

		r, err := NewReader(file)
		require.ErrorIs(t, err, ErrToReadReader)
		assert.Nil(t, r)
	})
	t.Run("Should return an error because reader is an invalid zip file", func(t *testing.T) {
		r, err := NewReader(os.Stdin)
		require.ErrorIs(t, err, zip.ErrFormat)
		assert.Nil(t, r)
	})
	t.Run("Should return a reader with a valid zip file", func(t *testing.T) {
		file, err := os.Open(testdata.TestZipFile)
		require.NoError(t, err)
		defer func() { require.NoError(t, file.Close()) }()

		r, err := NewReader(file)
		require.NoError(t, err)
		assert.NotNil(t, r)
	})
}

func TestReader_Write(t *testing.T) {
	t.Run("should return an error with a nil io.Writer", func(t *testing.T) {
		file, err := os.Open(testdata.TestZipFile)
		require.NoError(t, err)
		defer func() { require.NoError(t, file.Close()) }()

		r, err := NewReader(file)
		require.NoError(t, err)
		assert.NotNil(t, r)

		err = r.Write(nil, 0)
		require.ErrorIs(t, err, ErrNilWriter)
	})
	t.Run("should return an error with an invalid index", func(t *testing.T) {
		file, err := os.Open(testdata.TestZipFile)
		require.NoError(t, err)
		defer func() { require.NoError(t, file.Close()) }()

		r, err := NewReader(file)
		require.NoError(t, err)
		assert.NotNil(t, r)

		var buf bytes.Buffer
		err = r.Write(&buf, 1)
		require.ErrorIs(t, err, ErrInvalidIndex)
	})
	t.Run("should return an error to write data", func(t *testing.T) {
		var err error
		var file *os.File

		file, err = os.Open(testdata.TestZipFile)
		require.NoError(t, err)

		r, err := NewReader(file)
		require.NoError(t, err)

		err = r.Write(&failingWriter{}, 0)
		require.ErrorIs(t, err, ErrToWriteFile)
	})
	t.Run("should write the data file", func(t *testing.T) {
		var err error
		var file *os.File

		file, err = os.Open(testdata.TestZipFile)
		require.NoError(t, err)

		r, err := NewReader(file)
		require.NoError(t, err)

		var buf bytes.Buffer
		err = r.Write(&buf, 0)
		require.NoError(t, err)
		assert.Equal(t, testdata.ContentZipFile, buf.String())
	})
}

func TestReader_InfoFile(t *testing.T) {
	t.Run("should return an error with an invalid index", func(t *testing.T) {
		file, err := os.Open(testdata.TestZipFile)
		require.NoError(t, err)
		defer func() { require.NoError(t, file.Close()) }()

		r, err := NewReader(file)
		require.NoError(t, err)
		assert.NotNil(t, r)

		info, err := r.InfoFile(1)
		require.ErrorIs(t, err, ErrInvalidIndex)
		assert.Empty(t, info)
	})
	t.Run("Should return file info", func(t *testing.T) {
		file, err := os.Open(testdata.TestZipFile)
		require.NoError(t, err)
		defer func() { require.NoError(t, file.Close()) }()

		r, err := NewReader(file)
		require.NoError(t, err)
		assert.NotNil(t, r)

		info, err := r.InfoFile(0)
		require.NoError(t, err)
		assert.Equal(t, testdata.TestFileNameInZip, info.Name())
	})
}

func TestReader_Read(t *testing.T) {
	t.Run("should return an error with an invalid index", func(t *testing.T) {
		file, err := os.Open(testdata.TestZipFile)
		require.NoError(t, err)
		defer func() { require.NoError(t, file.Close()) }()

		r, err := NewReader(file)
		require.NoError(t, err)
		assert.NotNil(t, r)

		buf, err := r.Read(1)
		require.ErrorIs(t, err, ErrInvalidIndex)
		assert.Empty(t, buf)
	})
	t.Run("Should read the content zip file", func(t *testing.T) {
		file, err := os.Open(testdata.TestZipFile)
		require.NoError(t, err)
		defer func() { require.NoError(t, file.Close()) }()

		r, err := NewReader(file)
		require.NoError(t, err)
		assert.NotNil(t, r)

		buf, err := r.Read(0)
		require.NoError(t, err)
		assert.Equal(t, testdata.ContentZipFile, string(buf))
	})
}

func TestReader_NumFiles(t *testing.T) {
	t.Run("Should return the number of files in zip", func(t *testing.T) {
		file, err := os.Open(testdata.TestZipFile)
		require.NoError(t, err)
		defer func() { require.NoError(t, file.Close()) }()

		r, err := NewReader(file)
		require.NoError(t, err)
		assert.NotNil(t, r)

		assert.Equal(t, testdata.TestNumFile, r.NumFile())
	})
}
