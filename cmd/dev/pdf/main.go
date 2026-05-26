package main

import (
	"log"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var dummyText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."

func main() {
	cfg := config.NewBuilder().
		WithPageNumber().
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(buildHeader()...)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = m.RegisterFooter(buildFooter()...)
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(
		text.NewRow(20, "Main features", props.Text{Size: 15, Top: 6.5}),
	)
	m.AddRows(buildCodesRow()...)
	m.AddRows(buildImagesRow()...)
	m.AddRows(buildTextsRow()...)

	m.AddRows(
		text.NewRow(15, "Dummy Data", props.Text{Size: 12, Top: 5, Align: align.Center}),
	)

	for i := 0; i < 50; i++ {
		m.AddRows(text.NewRow(20, dummyText+dummyText+dummyText+dummyText+dummyText))
	}

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/v2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/v2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func buildCodesRow() []core.Row { _ = "STUB: not implemented"; return nil }

func buildImagesRow() []core.Row { _ = "STUB: not implemented"; return nil }

func buildTextsRow() []core.Row { _ = "STUB: not implemented"; return nil }

func buildHeader() []core.Row { _ = "STUB: not implemented"; return nil }

func buildFooter() []core.Row { _ = "STUB: not implemented"; return nil }
