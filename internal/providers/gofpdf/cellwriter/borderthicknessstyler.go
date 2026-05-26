package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type BorderThicknessStyler struct {
	stylerTemplate
	defaultLineThickness float64
}

func NewBorderThicknessStyler(fpdf gofpdfwrapper.Fpdf) *BorderThicknessStyler {
	_ = "STUB: not implemented"
	return nil
}

func (b *BorderThicknessStyler) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	_ = "STUB: not implemented"
	return
}
