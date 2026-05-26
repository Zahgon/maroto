// nolint: dupl
package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type FillColorStyler struct {
	stylerTemplate
	defaultFillColor *props.Color
}

func NewFillColorStyler(fpdf gofpdfwrapper.Fpdf) *FillColorStyler {
	_ = "STUB: not implemented"
	return nil
}

func (f *FillColorStyler) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	_ = "STUB: not implemented"
	return
}
