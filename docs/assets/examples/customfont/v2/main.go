// nolint:misspell // other languages are being classified as english misspells
package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"
)

func main() {
	m := GetMaroto("docs/assets/fonts/arial-unicode-ms.ttf")
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/customfontv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/customfontv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto(customFontFile string) core.Maroto {
	_ = "STUB: not implemented"
	return *new(core.Maroto)
}

func getLanguageSample() ([]string, [][]string) { _ = "STUB: not implemented"; return nil, nil }
