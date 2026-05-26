package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type CellWriter interface {
	SetNext(next CellWriter)
	GetNext() CellWriter
	GetName() string
	Apply(width, height float64, config *entity.Config, prop *props.Cell)
}

type cellWriter struct {
	stylerTemplate
	defaultColor *props.Color
}

func NewCellWriter(fpdf gofpdfwrapper.Fpdf) CellWriter {
	_ = "STUB: not implemented"
	return *new(CellWriter)
}

func (c *cellWriter) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	_ = "STUB: not implemented"
	return
}
