// Package config implements custom configuration builder.
// nolint:interfacebloat // there is no need to reduce this interface
package config

import (
	"time"

	"github.com/johnfercher/maroto/v2/pkg/consts/generation"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"

	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/protection"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/consts/provider"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// Builder is the abstraction responsible for global customizations on the document.
type Builder interface {
	WithPageSize(size pagesize.Type) Builder
	WithDimensions(width float64, height float64) Builder
	WithLeftMargin(left float64) Builder
	WithTopMargin(top float64) Builder
	WithRightMargin(right float64) Builder
	WithBottomMargin(bottom float64) Builder
	WithConcurrentMode(chunkWorkers int) Builder
	WithSequentialMode() Builder
	WithSequentialLowMemoryMode(chunkWorkers int) Builder
	WithDebug(on bool) Builder
	WithMaxGridSize(maxGridSize int) Builder
	WithDefaultFont(font *props.Font) Builder
	WithPageNumber(pageNumber ...props.PageNumber) Builder
	WithProtection(protectionType protection.Type, userPassword, ownerPassword string) Builder
	WithCompression(compression bool) Builder
	WithOrientation(orientation orientation.Type) Builder
	WithAuthor(author string, isUTF8 bool) Builder
	WithCreator(creator string, isUTF8 bool) Builder
	WithSubject(subject string, isUTF8 bool) Builder
	WithTitle(title string, isUTF8 bool) Builder
	WithCreationDate(time time.Time) Builder
	WithCustomFonts(customFonts []entity.CustomFont) Builder
	WithBackgroundImage(bytes []byte, extensionType extension.Type) Builder
	WithDisableAutoPageBreak(disabled bool) Builder
	WithKeywords(keywordsStr string, isUTF8 bool) Builder
	Build() *entity.Config
}

type CfgBuilder struct {
	providerType         provider.Type
	dimensions           *entity.Dimensions
	margins              *entity.Margins
	chunkWorkers         int
	debug                bool
	maxGridSize          int
	defaultFont          *props.Font
	customFonts          []entity.CustomFont
	pageNumber           *props.PageNumber
	protection           *entity.Protection
	compression          bool
	pageSize             *pagesize.Type
	orientation          orientation.Type
	metadata             *entity.Metadata
	backgroundImage      *entity.Image
	disableAutoPageBreak bool
	generationMode       generation.Mode
}

// NewBuilder is responsible to create an instance of Builder.
func NewBuilder() Builder { _ = "STUB: not implemented"; return *new(Builder) }

// WithPageSize defines the page size, ex: A4, A4 and etc.
func (b *CfgBuilder) WithPageSize(size pagesize.Type) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithDimensions defines custom page dimensions, this overrides page size.
func (b *CfgBuilder) WithDimensions(width float64, height float64) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithLeftMargin customize margin.
func (b *CfgBuilder) WithLeftMargin(left float64) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithTopMargin customize margin.
func (b *CfgBuilder) WithTopMargin(top float64) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithRightMargin customize margin.
func (b *CfgBuilder) WithRightMargin(right float64) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithBottomMargin customize margin.
func (b *CfgBuilder) WithBottomMargin(bottom float64) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithConcurrentMode defines concurrent generation, chunk workers define how mano chuncks
// will be executed concurrently.
func (b *CfgBuilder) WithConcurrentMode(chunkWorkers int) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithSequentialMode defines that maroto will run in default mode.
func (b *CfgBuilder) WithSequentialMode() Builder { _ = "STUB: not implemented"; return *new(Builder) }

// WithSequentialLowMemoryMode defines that maroto will run focusing in reduce memory consumption,
// chunk workers define how many divisions the work will have.
func (b *CfgBuilder) WithSequentialLowMemoryMode(chunkWorkers int) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithDebug defines a debug behaviour where maroto will draw borders in everything.
func (b *CfgBuilder) WithDebug(on bool) Builder {
	_ = "STUB: not implemented"
	return *

	// WithMaxGridSize defines a custom max grid sum which it will change the sum of column sizes.
	new(Builder)
}

func (b *CfgBuilder) WithMaxGridSize(maxGridSize int) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithDefaultFont defines a custom font, other than arial. This can be used to define a custom font as default.
func (b *CfgBuilder) WithDefaultFont(font *props.Font) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithPageNumber defines a string pattern to write the current page and total.
func (b *CfgBuilder) WithPageNumber(pageNumber ...props.PageNumber) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithProtection defines protection types to the PDF document.
func (b *CfgBuilder) WithProtection(protectionType protection.Type, userPassword, ownerPassword string) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithCompression defines compression.
func (b *CfgBuilder) WithCompression(compression bool) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithOrientation defines the page orientation. The default orientation is vertical,
// if horizontal is defined width and height will be flipped.
func (b *CfgBuilder) WithOrientation(orientation orientation.Type) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithAuthor defines the author name metadata.
func (b *CfgBuilder) WithAuthor(author string, isUTF8 bool) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithCreator defines the creator name metadata.
func (b *CfgBuilder) WithCreator(creator string, isUTF8 bool) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithSubject defines the subject metadata.
func (b *CfgBuilder) WithSubject(subject string, isUTF8 bool) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithTitle defines the title metadata.
func (b *CfgBuilder) WithTitle(title string, isUTF8 bool) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithCreationDate defines the creation date metadata.
func (b *CfgBuilder) WithCreationDate(time time.Time) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithCustomFonts add custom fonts.
func (b *CfgBuilder) WithCustomFonts(customFonts []entity.CustomFont) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithBackgroundImage defines the background image that will be applied in every page.
func (b *CfgBuilder) WithBackgroundImage(bytes []byte, ext extension.Type) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithDisableAutoPageBreak defines the option to disable automatic page breaks.
func (b *CfgBuilder) WithDisableAutoPageBreak(disabled bool) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// WithKeywords defines the document's keyword metadata.
func (b *CfgBuilder) WithKeywords(keywordsStr string, isUTF8 bool) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// Build finalizes the customization returning the entity.Config.
func (b *CfgBuilder) Build() *entity.Config { _ = "STUB: not implemented"; return nil }

func (b *CfgBuilder) getDimensions() *entity.Dimensions { _ = "STUB: not implemented"; return nil }
