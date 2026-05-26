// Package code implements creation of Barcode, MatrixCode and QrCode.
// nolint:dupl
package code

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type MatrixCode struct {
	code   string
	prop   props.Rect
	config *entity.Config
}

// NewMatrix is responsible to create an instance of a MatrixCode.
func NewMatrix(code string, barcodeProps ...props.Rect) core.Component {
	_ = "STUB: not implemented"
	return *new(core.Component)
}

// NewMatrixCol is responsible to create an instance of a MatrixCode wrapped in a Col.
func NewMatrixCol(size int, code string, ps ...props.Rect) core.Col {
	_ = "STUB: not implemented"
	return *new(core.Col)
}

// NewAutoMatrixRow is responsible to create an instance of a Matrix code wrapped in a Row with automatic height.
//   - code: The value that must be placed in the matrixcode
//   - ps: A set of settings that must be applied to the matrixcode
func NewAutoMatrixRow(code string, ps ...props.Rect) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// NewMatrixRow is responsible to create an instance of a MatrixCode wrapped in a Row.
func NewMatrixRow(height float64, code string, ps ...props.Rect) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// Render renders a MatrixCode into a PDF context.
func (m *MatrixCode) Render(provider core.Provider, cell *entity.Cell) {
	_ = "STUB: not implemented"
	return
}

// GetStructure returns the Structure of a MatrixCode.
func (m *MatrixCode) GetStructure() *node.Node[core.Structure] {
	_ = "STUB: not implemented"
	return nil
}

// GetHeight returns the height that the code will have in the PDF
func (m *MatrixCode) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	_ = "STUB: not implemented"
	return 0
}

// SetConfig sets the configuration of a MatrixCode.
func (m *MatrixCode) SetConfig(config *entity.Config) { _ = "STUB: not implemented"; return }
