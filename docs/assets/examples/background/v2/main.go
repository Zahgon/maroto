package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"
)

func main() {
	backgroundImage := "docs/assets/images/certificate.png"
	m := GetMaroto(backgroundImage)
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/backgroundv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/backgroundv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto(image string) core.Maroto { _ = "STUB: not implemented"; return *new(core.Maroto) }

func AddPage() core.Page { _ = "STUB: not implemented"; return *new(core.Page) }
