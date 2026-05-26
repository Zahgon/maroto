// Package entity contains all core entities.
package entity

// Cell represents a cell inside the PDF.
type Cell struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// NewRootCell creates the main Cell.
func NewRootCell(pageWidth, pageHeight float64, margins Margins) Cell {
	_ = "STUB: not implemented"
	return *new(Cell)
}

// GetDimensions returns the dimensions of the Cell (width and height).
func (c Cell) GetDimensions() *Dimensions { _ = "STUB: not implemented"; return nil }

// Copy deep copy the cell data.
func (c Cell) Copy() Cell { _ = "STUB: not implemented"; return *new(Cell) }
