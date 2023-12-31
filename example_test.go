package zip_test

import (
	"fmt"
	"log"
	"os"

	"github.com/gofast-pkg/zip"
)

const (
	testReadCloserZipFile = "./testdata/zipfile.zip"
)

func ExampleNew() {
	file, err := os.Open(testReadCloserZipFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r, err := zip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.NFiles())
	// Iterate through the files in the archive,
	for i := 0; i < r.NFiles(); i++ {
		var content []byte
		if content, err = r.ContentFile(i); err != nil {
			panic(err)
		}
		fmt.Println(string(content))
	}
	// Output:
	// 1
	//nolint:lll // buffer to match the example
	// annee_numero_de_tirage;1er_ou_2eme_tirage;jour_de_tirage;date_de_tirage;date_de_forclusion;boule_1;boule_2;boule_3;boule_4;boule_5;boule_6;boule_complementaire;combinaison_gagnante_en_ordre_croissant;numero_joker;nombre_de_gagnant_au_rang1;rapport_du_rang1;nombre_de_gagnant_au_rang2;rapport_du_rang2;nombre_de_gagnant_au_rang3;rapport_du_rang3;nombre_de_gagnant_au_rang4;rapport_du_rang4;nombre_de_gagnant_au_rang5;rapport_du_rang5;nombre_de_gagnant_au_rang6;rapport_du_rang6;nombre_de_gagnant_au_rang7;rapport_du_rang7;numero_jokerplus;devise;
	// 2008080;2;SA;20081004;20081204;33;32;42;16;15;49;37;15-16-32-33-42-49;;2;904940;8;10953,9;213;1400,3;589;61,6;11911;30,8;20552;5,4;255211;2,7;7 523 262;eur;
	// 2008080;1;SA;20081004;20081204;36;9;16;27;25;12;48;9-12-16-25-27-36;;3;282006;7;12513,4;461;662,5;1033;32;22997;16;25950;3,6;391755;1,8;7 523 262;eur;
}
