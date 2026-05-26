package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
)

// Signature represents properties from a signature.
type Signature struct {
	// FontFamily of the text, ex: consts.Arial, helvetica and etc.
	FontFamily string
	// FontStyle of the text, ex: consts.Normal, bold and etc.
	FontStyle fontstyle.Type
	// FontSize of the text.
	FontSize float64
	// FontColor define the font color.
	FontColor *Color
	// LineColor define the line color.
	LineColor *Color
	// LineStyle define the line style (solid or dashed).
	LineStyle linestyle.Type
	// LineThickness define the line thickness.
	LineThickness float64

	SafePadding float64
}

// ToMap returns a map with the Signature fields.
func (s *Signature) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }

// MakeValid from Signature define default values for a Signature.
func (s *Signature) MakeValid(defaultFontFamily string) { _ = "STUB: not implemented"; return }

// ToLineProp from Signature return a Line based on Signature.
func (s *Signature) ToLineProp(offsetPercent float64) *Line { _ = "STUB: not implemented"; return nil }

// ToFontProp from Signature return a Font based on Signature.
func (s *Signature) ToFontProp() *Font { _ = "STUB: not implemented"; return nil }

// ToTextProp from Signature return a Text based on Signature.
func (s *Signature) ToTextProp(align align.Type, top float64, verticalPadding float64) *Text {
	_ = "STUB: not implemented"
	return nil
}
