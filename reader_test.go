package zip

import (
	"archive/zip"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// files to have a zip.ReadCloser
const (
	testReadCloserZipFile = "testdata/zipfile.zip"
)

// files and folder use for the different tests case
const (
	testZipFile   = "test-file.zip"
	testFolder    = "./tmp"
	testFileInZip = "classic-loto-v1.csv"
	//nolint:lll // buffer for comparison
	testContentFileInZip = `annee_numero_de_tirage;1er_ou_2eme_tirage;jour_de_tirage;date_de_tirage;date_de_forclusion;boule_1;boule_2;boule_3;boule_4;boule_5;boule_6;boule_complementaire;combinaison_gagnante_en_ordre_croissant;numero_joker;nombre_de_gagnant_au_rang1;rapport_du_rang1;nombre_de_gagnant_au_rang2;rapport_du_rang2;nombre_de_gagnant_au_rang3;rapport_du_rang3;nombre_de_gagnant_au_rang4;rapport_du_rang4;nombre_de_gagnant_au_rang5;rapport_du_rang5;nombre_de_gagnant_au_rang6;rapport_du_rang6;nombre_de_gagnant_au_rang7;rapport_du_rang7;numero_jokerplus;devise;
2008080;2;SA;20081004;20081204;33;32;42;16;15;49;37;15-16-32-33-42-49;;2;904940;8;10953,9;213;1400,3;589;61,6;11911;30,8;20552;5,4;255211;2,7;7 523 262;eur;
2008080;1;SA;20081004;20081204;36;9;16;27;25;12;48;9-12-16-25-27-36;;3;282006;7;12513,4;461;662,5;1033;32;22997;16;25950;3,6;391755;1,8;7 523 262;eur;`
)

func includeIndexValidationTester(t *testing.T, f func(index int, r Reader) error) {
	t.Run("Should return an error with negative index", func(t *testing.T) {
		file, err := os.Open(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		r, err := NewReader(file)
		if err != nil {
			t.Error(err)
		}
		err = f(-1, r)
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrInvalidIndex)
		}
	})
	t.Run("Should return an error with over range index", func(t *testing.T) {
		file, err := os.Open(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		r, err := NewReader(file)
		if err != nil {
			t.Error(err)
		}
		err = f(math.MaxInt, r)
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrInvalidIndex)
		}
	})
}

func TestNewReader(t *testing.T) {
	t.Run("Should return an error because input is nil", func(t *testing.T) {
		r, err := NewReader(nil)

		if assert.Error(t, err) {
			assert.Nil(t, r)
			assert.ErrorIs(t, err, ErrInvalidInput)
		}
	})
	t.Run("Should return an error because io.ReadCloser is already closed", func(t *testing.T) {
		var err error
		var file *os.File

		expectedErr := "file already closed"
		if file, err = os.Open(testReadCloserZipFile); err != nil {
			t.Error(err)
		}
		if err = file.Close(); err != nil {
			t.Error(err)
		}

		r, err := NewReader(file)
		if assert.Error(t, err) {
			assert.Nil(t, r)
			assert.ErrorContains(t, err, expectedErr)
		}
	})
	t.Run("Should return an error because reader is an invalid zip file", func(t *testing.T) {
		r, err := NewReader(os.Stdin)

		if assert.Error(t, err) {
			assert.Nil(t, r)
			assert.ErrorIs(t, err, zip.ErrFormat)
		}
	})
	t.Run("Should return a reader with a valid zip file", func(t *testing.T) {
		file, err := os.Open(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		r, err := NewReader(file)

		if assert.NoError(t, err) {
			assert.NotNil(t, r)
		}
	})
}

func TestReader_WriteFile(t *testing.T) {
	t.Run("should return an error with index validation", func(t *testing.T) {
		includeIndexValidationTester(t, func(index int, r Reader) error {
			return r.WriteFile(index, testFolder)
		})
	})
	t.Run("Should return an error because folder path do not exist", func(t *testing.T) {
		var err error
		var file *os.File
		var r Reader

		if file, err = os.Open(testReadCloserZipFile); err != nil {
			t.Error(err)
		}
		defer file.Close()
		if r, err = NewReader(file); err != nil {
			t.Error(err)
		}

		if err = r.WriteFile(0, testFolder); err != nil {
			assert.Error(t, err)
			assert.ErrorIs(t, err, os.ErrNotExist)
		}
	})
	t.Run("Should write file", func(t *testing.T) {
		if err := os.Mkdir(testFolder, 0755); err != nil {
			t.Error(err)
		}
		defer os.RemoveAll(testFolder)
		file, err := os.Open(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		r, err := NewReader(file)
		if err != nil {
			t.Error(err)
		}

		err = r.WriteFile(0, testFolder)
		if assert.NoError(t, err) {
			assert.FileExists(t, testFolder+"/"+testFileInZip)
		}
	})
}

func TestReader_InfoFile(t *testing.T) {
	t.Run("should return an error with index validation", func(t *testing.T) {
		includeIndexValidationTester(t, func(index int, r Reader) error {
			_, err := r.InfoFile(index)

			return err
		})
	})
	t.Run("Should return file info", func(t *testing.T) {
		file, err := os.Open(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		r, err := NewReader(file)
		if err != nil {
			t.Error(err)
		}

		info, err := r.InfoFile(0)
		if assert.NoError(t, err) {
			assert.NotNil(t, info)
			assert.Equal(t, testFileInZip, info.Name())
		}
	})
}

func TestReader_ContentFile(t *testing.T) {
	t.Run("should return an error with index validation", func(t *testing.T) {
		includeIndexValidationTester(t, func(index int, r Reader) error {
			_, err := r.ContentFile(index)

			return err
		})
	})
	t.Run("Should return file content", func(t *testing.T) {
		file, err := os.Open(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		r, err := NewReader(file)
		if err != nil {
			t.Error(err)
		}

		content, err := r.ContentFile(0)
		if assert.NoError(t, err) {
			assert.Equal(t, string(content), string(testContentFileInZip))
		}
	})
}

func TestReader_NFiles(t *testing.T) {
	t.Run("Should return the number of files in zip", func(t *testing.T) {
		file, err := os.Open(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		r, err := NewReader(file)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 1, r.NFiles())
	})
}

func TestReader_Create(t *testing.T) {
	t.Run("Should return an error because file pathdo not exist", func(t *testing.T) {
		file, err := os.Open(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		r, err := NewReader(file)
		if err != nil {
			t.Error(err)
		}
		filepath := fmt.Sprintf("%s/%s", testFolder, testZipFile)

		err = r.Create(filepath)
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, os.ErrNotExist)
		}
	})
	t.Run("Should create a zip file", func(t *testing.T) {
		var err error
		var file *os.File
		var r Reader

		if err = os.Mkdir(testFolder, 0755); err != nil {
			t.Error(err)
		}
		defer os.RemoveAll(testFolder)

		if file, err = os.Open(testReadCloserZipFile); err != nil {
			t.Error(err)
		}
		defer file.Close()

		if r, err = NewReader(file); err != nil {
			t.Error(err)
		}
		filepath := fmt.Sprintf("%s/%s", testFolder, testZipFile)

		err = r.Create(filepath)
		if assert.NoError(t, err) {
			var content []byte

			// parse new zip file to check if it's valid
			file, err = os.Open(filepath)
			if assert.NoError(t, err) {
				defer file.Close()
			}
			r, err = NewReader(file)
			if assert.NoError(t, err) {
				assert.Equal(t, 1, r.NFiles())
			}
			content, err = r.ContentFile(0)
			if assert.NoError(t, err) {
				assert.Equal(t, string(content), string(testContentFileInZip))
			}
		}
	})
}
