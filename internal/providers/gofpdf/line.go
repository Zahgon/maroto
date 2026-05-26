package gofpdf

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Line struct {
	pdf              gofpdfwrapper.Fpdf
	defaultColor     *props.Color
	defaultThickness float64
}

func NewLine(pdf gofpdfwrapper.Fpdf) *Line { _ = "STUB: not implemented"; return nil }

func (l *Line) Add(cell *entity.Cell, prop *props.Line) { _ = "STUB: not implemented"; return }

func (l *Line) renderVertical(cell *entity.Cell, prop *props.Line) {
	_ = "STUB: not implemented"
	return
}

func (l *Line) renderHorizontal(cell *entity.Cell, prop *props.Line) {
	_ = "STUB: not implemented"
	return
}
