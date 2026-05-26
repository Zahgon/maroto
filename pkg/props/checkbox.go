package props

// Checkbox represents properties from a Checkbox inside a cell.
type Checkbox struct {
	// Checked defines whether the checkbox is marked.
	Checked bool
	// Top is the amount of space between the upper cell limit and the checkbox.
	Top float64
	// Left is the amount of space between the left cell boundary and the checkbox.
	Left float64
	// Size is the size of the checkbox square in mm.
	Size float64
}

// ToMap converts a Checkbox to a map.
func (c *Checkbox) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }

// MakeValid from Checkbox define default values for a Checkbox.
func (c *Checkbox) MakeValid() { _ = "STUB: not implemented"; return }
