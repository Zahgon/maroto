// Package code implements creation of Barcode, MatrixCode and QrCode.
// nolint:dupl // It's similar to Barcode.go and it's hard to extract common code.
package code

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Barcode struct {
	code   string
	prop   props.Barcode
	config *entity.Config
}

// NewBar is responsible to create an instance of a Barcode.
//   - code: The value that must be placed in the barcode
//   - ps: A set of settings that must be applied to the barcode
func NewBar(code string, ps ...props.Barcode) core.Component {
	_ = "STUB: not implemented"
	return *new(core.Component)
}

// NewBarCol is responsible to create an instance of a Barcode wrapped in a Col.
//   - size: O tamanho da coluna
//   - code: The value that must be placed in the barcode
//   - ps: A set of settings that must be applied to the barcode
func NewBarCol(size int, code string, ps ...props.Barcode) core.Col {
	_ = "STUB: not implemented"
	return *new(core.Col)
}

// NewBarRow is responsible to create an instance of a Barcode wrapped in a Row.
// using this method the col size will be automatically set to the maximum value
//   - height: The height of the line
//   - code: The value that must be placed in the barcode
//   - ps: A set of settings that must be applied to the barcode
func NewBarRow(height float64, code string, ps ...props.Barcode) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// NewAutoBarRow is responsible to create an instance of a Barcode wrapped in a Row with automatic height.
// using this method the col size will be automatically set to the maximum value
//   - code: The value that must be placed in the barcode
//   - ps: A set of settings that must be applied to the barcode
func NewAutoBarRow(code string, ps ...props.Barcode) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// Render renders a Barcode into a PDF context. The maroto cal this method in process to
// generate the pdf.
//   - provider: Is the creator provider used to generate the pdf
//   - cell: cell represents the space available to draw the component
func (b *Barcode) Render(provider core.Provider, cell *entity.Cell) {
	_ = "STUB: not implemented"
	return
}

// GetStructure returns the structure of a barcode. This method is typically used when creating tests
func (b *Barcode) GetStructure() *node.Node[core.Structure] { _ = "STUB: not implemented"; return nil }

// GetHeight returns the height that the barcode will have in the PDF
func (b *Barcode) GetHeight(_ core.Provider, cell *entity.Cell) float64 {
	_ = "STUB: not implemented"
	return 0
}

// SetConfig sets the configuration of a Barcode.
func (b *Barcode) SetConfig(config *entity.Config) { _ = "STUB: not implemented"; return }
