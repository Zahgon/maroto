package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
)

// Line represents properties from a Line inside a cell.
type Line struct {
	// Color define the line color.
	Color *Color
	// Style define the line style (solid or dashed).
	Style linestyle.Type
	// Thickness define the line thicknesl.
	Thickness float64
	// Orientation define if line would be horizontal or vertical.
	Orientation orientation.Type
	// OffsetPercent define where the line would be placed, 0 is the start of cell, 50 the middle and 100 the end.
	OffsetPercent float64
	// SizePercent define the size of the line inside cell.
	SizePercent float64
}

// ToMap returns a map with the Line fields.
func (l *Line) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }

// MakeValid from Line define default values for a Line.
func (l *Line) MakeValid() { _ = "STUB: not implemented"; return }
