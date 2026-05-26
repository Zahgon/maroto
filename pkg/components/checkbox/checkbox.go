// Package checkbox implements creation of checkboxes.
package checkbox

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Checkbox struct {
	label  string
	prop   props.Checkbox
	config *entity.Config
}

// New is responsible to create an instance of a Checkbox.
func New(label string, ps ...props.Checkbox) core.Component {
	_ = "STUB: not implemented"
	return *new(core.Component)
}

// NewCol is responsible to create an instance of a Checkbox wrapped in a Col.
func NewCol(size int, label string, ps ...props.Checkbox) core.Col {
	_ = "STUB: not implemented"
	return *new(core.Col)
}

// NewAutoRow is responsible for creating an instance of Checkbox grouped in a Row with automatic height.
func NewAutoRow(label string, ps ...props.Checkbox) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// NewRow is responsible to create an instance of a Checkbox wrapped in a Row.
func NewRow(height float64, label string, ps ...props.Checkbox) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// GetStructure returns the Structure of a Checkbox.
func (c *Checkbox) GetStructure() *node.Node[core.Structure] { _ = "STUB: not implemented"; return nil }

// SetConfig sets the config.
func (c *Checkbox) SetConfig(config *entity.Config) {
	_ = "STUB: not implemented"

	// GetHeight returns the height that the checkbox will have in the PDF.
	return
}

func (c *Checkbox) GetHeight(_ core.Provider, _ *entity.Cell) float64 {
	_ = "STUB: not implemented"
	return 0
}

// Render renders a Checkbox into a PDF context.
func (c *Checkbox) Render(provider core.Provider, cell *entity.Cell) {
	_ = "STUB: not implemented"
	return
}
