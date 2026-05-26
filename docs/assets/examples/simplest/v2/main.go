package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	err = document.Save("docs/assets/pdf/simplestv2.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

func GetMaroto() core.Maroto { _ = "STUB: not implemented"; return *new(core.Maroto) }
