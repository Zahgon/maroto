// Package col implements creation of columns.
package col

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Col struct {
	size       int
	isMax      bool
	components []core.Component
	config     *entity.Config
	style      *props.Cell
}

// New is responsible to create an instance of core.Col.
func New(size ...int) core.Col { _ = "STUB: not implemented"; return *new(core.Col) }

// Add is responsible to add a component to a core.Col.
func (c *Col) Add(components ...core.Component) core.Col {
	_ = "STUB: not implemented"
	return *new(core.Col)
}

// GetSize returns the size of a core.Col.
func (c *Col) GetSize() int { _ = "STUB: not implemented"; return 0 }

// GetStructure returns the Structure of a core.Col.
func (c *Col) GetStructure() *node.Node[core.Structure] { _ = "STUB: not implemented"; return nil }

// Render renders a core.Col into a PDF context.
func (c *Col) Render(provider core.Provider, cell entity.Cell, createCell bool) {
	_ = "STUB: not implemented"
	return
}

// SetConfig set the config for the component.
func (c *Col) SetConfig(config *entity.Config) { _ = "STUB: not implemented"; return }

// WithStyle sets the style for the column.
func (c *Col) WithStyle(style *props.Cell) core.Col {
	_ = "STUB: not implemented"
	return *new(core.Col)
}

// GetHeight returns the height of the column content
func (c *Col) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	_ = "STUB: not implemented"
	return 0
}
