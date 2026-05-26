// nolint:misspell // other languages are being classified as english typos
package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"
)

func main() {
	m := GetMaroto("docs/assets/images/frontpage.png")
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/compressionv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/compressionv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto(imagePath string) core.Maroto { _ = "STUB: not implemented"; return *new(core.Maroto) }
