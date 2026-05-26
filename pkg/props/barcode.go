package props

import "github.com/johnfercher/maroto/v2/pkg/consts/barcode"

// Barcode represents properties from a barcode inside a cell.
type Barcode struct {
	// Left is the space between the left cell boundary to the barcode, if center is false.
	Left float64
	// Top is space between the upper cell limit to the barcode, if center is false.
	Top float64
	// Percent is how much the barcode will occupy the cell,
	// ex 100%: The barcode will fulfill the entire cell
	// ex 50%: The greater side from the barcode will have half the size of the cell.
	Percent float64
	// Proportion is the proportion between size of the barcode.
	// Ex: 16x9, 4x3...
	Proportion Proportion
	// Center define that the barcode will be vertically and horizontally centralized.
	Center bool
	// Type represents the barcode type. Default: code128
	Type barcode.Type
}

// ToMap from Barcode will return a map representation from Barcode.
func (b *Barcode) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }

// ToRectProp from Barcode will return a Rect representation from Barcode.
func (b *Barcode) ToRectProp() *Rect { _ = "STUB: not implemented"; return nil }

// MakeValid from Barcode will make the properties from a barcode reliable to fit inside a cell
// and define default values for a barcode.
func (b *Barcode) MakeValid() { _ = "STUB: not implemented"; return }
