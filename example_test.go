package zip_test

import (
	"errors"
	"fmt"
	"os"

	"github.com/gofast-pkg/zip"
	"github.com/gofast-pkg/zip/testdata"
)

func ExampleNewReader() {
	file, err := os.Open(testdata.TestZipFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = errors.Join(err, file.Close()); err != nil {
			panic(err)
		}
	}()

	r, err := zip.NewReader(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.NumFile())

	// Iterate through the files in the archive,
	for i := 0; i < r.NumFile(); i++ {
		var content []byte
		if content, err = r.Read(i); err != nil {
			panic(err)
		}

		fmt.Println(string(content[:5]))
		fmt.Println(string(content[len(content)-5:]))
	}
	// Output:
	// 1
	// annee
	// ;eur;
}
