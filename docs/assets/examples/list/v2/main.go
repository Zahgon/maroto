package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var background = &props.Color{
	Red:   200,
	Green: 200,
	Blue:  200,
}

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/listv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/listv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto { _ = "STUB: not implemented"; return *new(core.Maroto) }

type Object struct {
	Key   string
	Value string
}

func (o Object) GetHeader() core.Row { _ = "STUB: not implemented"; return *new(core.Row) }

func (o Object) GetContent(i int) core.Row { _ = "STUB: not implemented"; return *new(core.Row) }

func getObjects(max int) []Object { _ = "STUB: not implemented"; return nil }
