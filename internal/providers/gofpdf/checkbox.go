package gofpdf

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

const labelGap = 1.0

type Checkbox struct {
	pdf  gofpdfwrapper.Fpdf
	font core.Font
}

// NewCheckbox create a Checkbox.
func NewCheckbox(pdf gofpdfwrapper.Fpdf, font core.Font) *Checkbox {
	_ = "STUB: not implemented"
	return nil
}

// Add a checkbox with a label inside a cell.
func (c *Checkbox) Add(label string, cell *entity.Cell, prop *props.Checkbox) {
	_ = "STUB: not implemented"
	return
}

// Draw the checkbox square border

// Draw X mark inside the box

// Draw label to the right of the checkbox, vertically centered
