package gofpdf

import (
	"github.com/johnfercher/maroto/v2/internal/cache"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

// Dependencies is the dependencies provider for gofpdf.
type Dependencies struct {
	Fpdf       gofpdfwrapper.Fpdf
	Font       core.Font
	Text       core.Text
	Code       core.Code
	Image      core.Image
	Line       core.Line
	Checkbox   core.Checkbox
	Cache      cache.Cache
	CellWriter cellwriter.CellWriter
	Cfg        *entity.Config
}

// Builder is the dependencies builder for gofpdf.
type Builder interface {
	Build(cfg *entity.Config, cache cache.Cache) *Dependencies
}

type builder struct{}

// NewBuilder create a new Builder
func NewBuilder() Builder {
	_ = "STUB: not implemented"

	// Build create a new Dependencies.
	return *new(Builder)
}

func (b *builder) Build(cfg *entity.Config, cache cache.Cache) *Dependencies {
	_ = "STUB: not implemented"
	return nil
}
