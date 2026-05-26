// nolint: dupl
package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type BorderColorStyler struct {
	stylerTemplate
	defaultColor *props.Color
}

func NewBorderColorStyler(fpdf gofpdfwrapper.Fpdf) *BorderColorStyler {
	_ = "STUB: not implemented"
	return nil
}

func (b *BorderColorStyler) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	_ = "STUB: not implemented"
	return
}
