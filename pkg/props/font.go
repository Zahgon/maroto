package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

// Font represents properties from a text.
type Font struct {
	// Family of the text, ex: constf.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: constf.Normal, bold and etc.
	Style fontstyle.Type
	// Size of the text.
	Size float64
	// Color define the font color.
	Color *Color
}

// AppendMap appends the font fields to a map.
func (f *Font) AppendMap(m map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

// MakeValid from Font define default values for a Signature.
func (f *Font) MakeValid(defaultFamily string) { _ = "STUB: not implemented"; return }

// ToTextProp from Font return a Text based on Font.
func (f *Font) ToTextProp(align align.Type, top float64, verticalPadding float64) *Text {
	_ = "STUB: not implemented"
	return nil
}
