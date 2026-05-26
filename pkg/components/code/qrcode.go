// Package code implements creation of Barcode, MatrixCode and QrCode.
// nolint:dupl
package code

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type QrCode struct {
	code   string
	prop   props.Rect
	config *entity.Config
}

// NewQr is responsible to create an instance of a QrCode.
func NewQr(code string, barcodeProps ...props.Rect) core.Component {
	_ = "STUB: not implemented"
	return *new(core.Component)
}

// NewQrCol is responsible to create an instance of a QrCode wrapped in a Col.
func NewQrCol(size int, code string, ps ...props.Rect) core.Col {
	_ = "STUB: not implemented"
	return *new(core.Col)
}

// NewAutoMatrixRow is responsible to create an instance of a qrcode wrapped in a Row with automatic height.
//   - code: The value that must be placed in the qrcode
//   - ps: A set of settings that must be applied to the qrcode
func NewAutoQrRow(code string, ps ...props.Rect) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// NewQrRow is responsible to create an instance of a QrCode wrapped in a Row.
func NewQrRow(height float64, code string, ps ...props.Rect) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// Render renders a QrCode into a PDF context.
func (q *QrCode) Render(provider core.Provider, cell *entity.Cell) {
	_ = "STUB: not implemented"
	return
}

// GetStructure returns the Structure of a QrCode.
func (q *QrCode) GetStructure() *node.Node[core.Structure] { _ = "STUB: not implemented"; return nil }

// GetHeight returns the height that the QrCode will have in the PDF
func (q *QrCode) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	_ = "STUB: not implemented"
	return 0
}

// SetConfig set the config for the component.
func (q *QrCode) SetConfig(config *entity.Config) { _ = "STUB: not implemented"; return }
