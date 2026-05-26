package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/metrics"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var dummyText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."

var background = &props.Color{
	Red:   200,
	Green: 200,
	Blue:  200,
}

func main() {
	var builder strings.Builder
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		builder.WriteString(fmt.Sprintf("%f", run().Value) + "\n")
	}

	err := os.WriteFile("docs/assets/text/benchmark.txt", []byte(builder.String()), os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func run() *metrics.Time { _ = "STUB: not implemented"; return nil }

func buildCodesRow() []core.Row { _ = "STUB: not implemented"; return nil }

func buildImagesRow() []core.Row { _ = "STUB: not implemented"; return nil }

func buildTextsRow() []core.Row { _ = "STUB: not implemented"; return nil }

func buildHeader() []core.Row { _ = "STUB: not implemented"; return nil }

func buildFooter() []core.Row { _ = "STUB: not implemented"; return nil }

type Object struct {
	Key   string
	Value string
}

func (o Object) GetHeader() core.Row { _ = "STUB: not implemented"; return *new(core.Row) }

func (o Object) GetContent(i int) core.Row { _ = "STUB: not implemented"; return *new(core.Row) }

func getObjects(maxObjects int) []Object { _ = "STUB: not implemented"; return nil }
