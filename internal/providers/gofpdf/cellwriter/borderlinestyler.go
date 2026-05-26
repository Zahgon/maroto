package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type BorderLineStyler struct {
	stylerTemplate
}

func NewBorderLineStyler(fpdf gofpdfwrapper.Fpdf) *BorderLineStyler {
	_ = "STUB: not implemented"
	return nil
}

func (b *BorderLineStyler) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	_ = "STUB: not implemented"
	return
}
