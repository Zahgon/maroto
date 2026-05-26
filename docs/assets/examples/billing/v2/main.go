package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/billingv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/billingv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto { _ = "STUB: not implemented"; return *new(core.Maroto) }

func getTransactions() []core.Row { _ = "STUB: not implemented"; return nil }

func getPageHeader() core.Row { _ = "STUB: not implemented"; return *new(core.Row) }

func getPageFooter() core.Row { _ = "STUB: not implemented"; return *new(core.Row) }

func getDarkGrayColor() *props.Color { _ = "STUB: not implemented"; return nil }

func getGrayColor() *props.Color { _ = "STUB: not implemented"; return nil }

func getBlueColor() *props.Color { _ = "STUB: not implemented"; return nil }

func getRedColor() *props.Color { _ = "STUB: not implemented"; return nil }

func getContents() [][]string { _ = "STUB: not implemented"; return nil }
