// Package border contains all border types.
package border

// Type represents a border type.
type Type int

// None is the default border type.
const None Type = 0

const (
	// Left is a border type that borders the left side.
	Left Type = 1 << iota
	// Top is a border type that borders the top side.
	Top
	// Right is a border type that borders the right side.
	Right
	// Bottom is a border type that borders the bottom side.
	Bottom
	// Full is a border type that borders all sides.
	Full = Left | Top | Right | Bottom
)

// IsValid checks if the border type is valid.
func (t Type) IsValid() bool { _ = "STUB: not implemented"; return false }

// HasLeft checks if the border type includes left border.
func (t Type) HasLeft() bool {
	_ = "STUB: not implemented"

	// HasTop checks if the border type includes top border.
	return false
}

func (t Type) HasTop() bool {
	_ = "STUB: not implemented"

	// HasRight checks if the border type includes right border.
	return false
}

func (t Type) HasRight() bool { _ = "STUB: not implemented"; return false }

// HasBottom checks if the border type includes bottom border.
func (t Type) HasBottom() bool { _ = "STUB: not implemented"; return false }

// String returns the string representation of the border type.
func (t Type) String() string { _ = "STUB: not implemented"; return "" }
