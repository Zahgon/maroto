package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

// Place is the representation of a place in a page.
type Place string

const (
	// LeftTop is the place in the left top of the page.
	LeftTop Place = "left_top"
	// Top is the place in the top of the page.
	Top Place = "top"
	// RightTop is the place in the right top of the page.
	RightTop Place = "right_top"
	// LeftBottom is the place in the left bottom of the page.
	LeftBottom Place = "left_bottom"
	// Bottom is the place in the bottom of the page.
	Bottom Place = "bottom"
	// RightBottom is the place in the right bottom of the page.
	RightBottom Place = "right_bottom"
)

// IsValid checks if the place is valid.
func (p Place) IsValid() bool { _ = "STUB: not implemented"; return false }

// PageNumber have attributes of page number.
type PageNumber struct {
	// Pattern is the string pattern which will be used to apply the page count component.
	Pattern string
	// Place defines where the page count component will be placed.
	Place Place
	// Family defines which font family will be applied to page count.
	Family string
	// Style defines which font style will be applied to page count.
	Style fontstyle.Type
	// Size defines which font size will be applied to page count.
	Size float64
	// Color defines which will be applied to page count.
	Color *Color
}

// GetNumberTextProp returns the Text properties of the page number.
// nolint:staticcheck // builder
func (p *PageNumber) GetNumberTextProp(height float64) *Text { _ = "STUB: not implemented"; return nil }

// GetPageString returns the page string.
func (p *PageNumber) GetPageString(current, total int) string { _ = "STUB: not implemented"; return "" }

// WithFont apply font if not defined before.
func (p *PageNumber) WithFont(font *Font) { _ = "STUB: not implemented"; return }

// AppendMap appends the font fields to a map.
func (p *PageNumber) AppendMap(m map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
