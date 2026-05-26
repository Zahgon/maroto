// Package text implements creation of texts.
package text

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Text struct {
	value  string
	prop   props.Text
	config *entity.Config
}

// New is responsible to create an instance of a Text.
func New(value string, ps ...props.Text) core.Component {
	_ = "STUB: not implemented"
	return *new(core.Component)
}

// NewCol is responsible to create an instance of a Text wrapped in a Col.
func NewCol(size int, value string, ps ...props.Text) core.Col {
	_ = "STUB: not implemented"
	return *new(core.Col)
}

// NewAutoRow is responsible for creating an instance of Text grouped in a Line with automatic height.
func NewAutoRow(value string, ps ...props.Text) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// NewRow is responsible to create an instance of a Text wrapped in a Row.
func NewRow(height float64, value string, ps ...props.Text) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// GetStructure returns the Structure of a Text.
func (t *Text) GetStructure() *node.Node[core.Structure] { _ = "STUB: not implemented"; return nil }

// GetHeight returns the height that the text will have in the PDF
func (t *Text) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	_ = "STUB: not implemented"
	return 0
}

// SetConfig sets the config.
func (t *Text) SetConfig(config *entity.Config) { _ = "STUB: not implemented"; return }

// Render renders a Text into a PDF context.
func (t *Text) Render(provider core.Provider, cell *entity.Cell) { _ = "STUB: not implemented"; return }
