// Package row implements creation of rows.
package row

import (
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Row struct {
	height     float64
	autoHeight bool
	cols       []core.Col
	style      *props.Cell
	config     *entity.Config
}

// New is responsible to create a core.Row.
//
// Height is an optional parameter that, if not sent, will be calculated automatically
// Height is defined in mm.
func New(height ...float64) core.Row { _ = "STUB: not implemented"; return *new(core.Row) }

// SetConfig sets the Row configuration.
func (r *Row) SetConfig(config *entity.Config) { _ = "STUB: not implemented"; return }

// Add is responsible to add one or more core.Col to a core.Row.
func (r *Row) Add(cols ...core.Col) core.Row { _ = "STUB: not implemented"; return *new(core.Row) }

// GetColumns returns the columns of a core.Row.
func (r *Row) GetColumns() []core.Col {
	_ = "STUB: not implemented"

	// GetHeight returns the height of a core.Row.
	return nil
}

func (r *Row) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	_ = "STUB: not implemented"
	return 0
}

// GetStructure returns the Structure of a core.Row.
func (r *Row) GetStructure() *node.Node[core.Structure] { _ = "STUB: not implemented"; return nil }

// Render renders a Row into a PDF context.
func (r *Row) Render(provider core.Provider, cell entity.Cell) { _ = "STUB: not implemented"; return }

// WithStyle sets the style of a Row.
func (r *Row) WithStyle(style *props.Cell) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// resetHeight resets the line height to 0
func (r *Row) resetHeight() {
	_ = "STUB: not implemented"

	// Returns the height of the row content
	return
}

func (r *Row) getBiggestCol(provider core.Provider, cell *entity.Cell) float64 {
	_ = "STUB: not implemented"
	return 0
}
