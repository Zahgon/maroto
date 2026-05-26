package gofpdf

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Text struct {
	pdf  gofpdfwrapper.Fpdf
	math core.Math
	font core.Font
}

// NewText create a Text.
func NewText(pdf gofpdfwrapper.Fpdf, math core.Math, font core.Font) *Text {
	_ = "STUB: not implemented"
	return nil
}

// Add a text inside a cell.
func (s *Text) Add(text string, cell *entity.Cell, textProp *props.Text) {
	_ = "STUB: not implemented"
	return
}

// override style if hyperlink is set

// Apply Unicode before calc spaces

// If should add one line

// GetLinesQuantity retrieve the quantity of lines which a text will occupy to avoid that text to extrapolate a cell.
func (s *Text) GetLinesQuantity(text string, textProp *props.Text, colWidth float64) int {
	_ = "STUB: not implemented"
	return 0
}

func (s *Text) getLinesBreakingLineFromSpace(words []string, colWidth float64) []string {
	_ = "STUB: not implemented"
	return nil
}

func (s *Text) getLinesBreakingLineWithDash(words string, colWidth float64) []string {
	_ = "STUB: not implemented"
	return nil
}

func (s *Text) addLine(textProp *props.Text, xColOffset, colWidth, yColOffset, textWidth float64, text string) {
	_ = "STUB: not implemented"
	return
}

func (s *Text) textToUnicode(txt string, props *props.Text) string {
	_ = "STUB: not implemented"
	return ""
}

func isIncorrectSpaceWidth(textWidth, spaceWidth, defaultSpaceWidth float64, text string) bool {
	_ = "STUB: not implemented"
	return false
}
