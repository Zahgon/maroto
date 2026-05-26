package gofpdf

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

const (
	gofpdfFontScale1 = 72.0
	gofpdfFontScale2 = 25.4
)

type Font struct {
	pdf         gofpdfwrapper.Fpdf
	size        float64
	family      string
	style       fontstyle.Type
	scaleFactor float64
	fontColor   *props.Color
}

// NewFont create a Font.
func NewFont(pdf gofpdfwrapper.Fpdf, size float64, family string, style fontstyle.Type) *Font {
	_ = "STUB: not implemented"
	return nil
}

// Bytes defined inside gofpdf constructor,

// GetFamily return the currently Font family configured.
func (s *Font) GetFamily() string {
	_ = "STUB: not implemented"

	// GetStyle return the currently Font style configured.
	return ""
}

func (s *Font) GetStyle() fontstyle.Type {
	_ = "STUB: not implemented"

	// GetSize return the currently Font size configured.
	return *new(fontstyle.Type)
}

func (s *Font) GetSize() float64 {
	_ = "STUB: not implemented"

	// GetFont return all the currently Font properties configured.
	return 0
}

func (s *Font) GetFont() (string, fontstyle.Type, float64) {
	_ = "STUB: not implemented"
	return "", *new(fontstyle.Type), 0
}

func (s *Font) GetHeight(family string, style fontstyle.Type, size float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

// SetFamily defines a new Font family.
func (s *Font) SetFamily(family string) { _ = "STUB: not implemented"; return }

// SetStyle defines a new Font style.
func (s *Font) SetStyle(style fontstyle.Type) { _ = "STUB: not implemented"; return }

// SetSize defines a new Font size.
func (s *Font) SetSize(size float64) { _ = "STUB: not implemented"; return }

// SetFont defines all new Font properties.
func (s *Font) SetFont(family string, style fontstyle.Type, size float64) {
	_ = "STUB: not implemented"
	return
}

func (s *Font) SetColor(color *props.Color) { _ = "STUB: not implemented"; return }

func (s *Font) GetColor() *props.Color { _ = "STUB: not implemented"; return nil }
