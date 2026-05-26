package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type stylerTemplate struct {
	next CellWriter
	fpdf gofpdfwrapper.Fpdf
	name string
}

func (s *stylerTemplate) SetNext(next CellWriter) { _ = "STUB: not implemented"; return }

func (s *stylerTemplate) GetName() string { _ = "STUB: not implemented"; return "" }

func (s *stylerTemplate) GetNext() CellWriter { _ = "STUB: not implemented"; return *new(CellWriter) }

func (s *stylerTemplate) GoToNext(width, height float64, config *entity.Config, prop *props.Cell) {
	_ = "STUB: not implemented"
	return
}
