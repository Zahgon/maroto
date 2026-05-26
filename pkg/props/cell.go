package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
)

// Cell is the representation of a cell in the grid system.
// This can be applied to Col or Row.
type Cell struct {
	// BackgroundColor defines which color will be applied to a cell.
	// Default: nil
	BackgroundColor *Color
	// BorderColor defines which color will be applied to a border cell
	// Default: nil
	BorderColor *Color
	// BorderType defines which kind of border will be applied to a cell.
	// Default: border.None
	BorderType border.Type
	// BorderThickness defines the border thickness applied to a cell.
	// Default: 0.2
	BorderThickness float64
	// LineStyle defines which line style will be applied to a cell.
	// Default: Solid
	LineStyle linestyle.Type
}

// ToMap adds the Cell fields to the map.
func (c *Cell) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }
