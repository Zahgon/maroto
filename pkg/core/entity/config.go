package entity

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/generation"
	"github.com/johnfercher/maroto/v2/pkg/consts/provider"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// Config is the configuration of a maroto instance.
type Config struct {
	ProviderType         provider.Type
	Dimensions           *Dimensions
	Margins              *Margins
	DefaultFont          *props.Font
	CustomFonts          []CustomFont
	GenerationMode       generation.Mode
	ChunkWorkers         int
	Debug                bool
	MaxGridSize          int
	PageNumber           *props.PageNumber
	Protection           *Protection
	Compression          bool
	Metadata             *Metadata
	BackgroundImage      *Image
	DisableAutoPageBreak bool
}

// ToMap converts Config to a map[string]any .
func (c *Config) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }
