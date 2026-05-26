// Package image implements creation of images from file and bytes.
package image

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type FileImage struct {
	path   string
	prop   props.Rect
	config *entity.Config
}

// NewFromFile is responsible to create an instance of an Image.
func NewFromFile(path string, ps ...props.Rect) core.Component {
	_ = "STUB: not implemented"
	return *new(core.Component)
}

// NewFromFileCol is responsible to create an instance of an Image wrapped in a Col.
func NewFromFileCol(size int, path string, ps ...props.Rect) core.Col {
	_ = "STUB: not implemented"
	return *new(core.Col)
}

// NewFromFileRow is responsible to create an instance of an Image wrapped in a Row.
func NewFromFileRow(height float64, path string, ps ...props.Rect) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// NewFromFileRow is responsible to create an instance of an Image wrapped in a automatic Row.
func NewAutoFromFileRow(path string, ps ...props.Rect) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// Render renders an Image into a PDF context.
func (f *FileImage) Render(provider core.Provider, cell *entity.Cell) {
	_ = "STUB: not implemented"
	return
}

// GetStructure returns the Structure of an Image.
func (f *FileImage) GetStructure() *node.Node[core.Structure] {
	_ = "STUB: not implemented"
	return nil
}

// GetHeight returns the height that the image will have in the PDF
func (f *FileImage) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	_ = "STUB: not implemented"
	return 0
}

// SetConfig sets the pdf config.
func (f *FileImage) SetConfig(config *entity.Config) { _ = "STUB: not implemented"; return }
