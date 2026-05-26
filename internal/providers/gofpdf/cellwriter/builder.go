package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
)

type WriterBuilder struct{}

func NewBuilder() *WriterBuilder { _ = "STUB: not implemented"; return nil }

func (c *WriterBuilder) Build(fpdf gofpdfwrapper.Fpdf) CellWriter {
	_ = "STUB: not implemented"
	return *new(CellWriter)
}
