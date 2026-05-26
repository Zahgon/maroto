package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

// Text represents properties from a Text inside a cell.
type Text struct {
	// Top is the amount of space between the upper cell limit and the text.
	Top float64
	// Bottom is the amount of space between the lower cell limit and the text. (Used by auto row only)
	Bottom float64
	// Left is the minimal amount of space between the left cell boundary and the text.
	Left float64
	// Right is the minimal amount of space between the right cell boundary and the text.
	Right float64
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style fontstyle.Type
	// Size of the text.
	Size float64
	// Align of the text.
	Align align.Type
	// BreakLineStrategy define the break line strategy.
	BreakLineStrategy breakline.Strategy
	// VerticalPadding define an additional space between linet.
	VerticalPadding float64
	// Color define the font style color.
	Color *Color
	// Hyperlink define a link to be opened when the text is clicked.
	Hyperlink *string
}

// ToMap converts a Text to a map.
func (t *Text) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }

// MakeValid from Text define default values for a Text.
func (t *Text) MakeValid(font *Font) { _ = "STUB: not implemented"; return }
