package fixture

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// TextProp is responsible to give a valid props.Text.
func TextProp() props.Text { _ = "STUB: not implemented"; return *new(props.Text) }

// FontProp is responsible to give a valid props.Font.
func FontProp() props.Font { _ = "STUB: not implemented"; return *new(props.Font) }

// BarcodeProp is responsible to give a valid props.Barcode.
func BarcodeProp() props.Barcode { _ = "STUB: not implemented"; return *new(props.Barcode) }

// RectProp is responsible to give a valid props.Rect.
func RectProp() props.Rect { _ = "STUB: not implemented"; return *new(props.Rect) }

// ConfigEntity is responsible to give a valid entity.Config.
func ConfigEntity() entity.Config { _ = "STUB: not implemented"; return *new(entity.Config) }

// CellEntity is responsible to give a valid entity.Cell.
func CellEntity() entity.Cell { _ = "STUB: not implemented"; return *new(entity.Cell) }

// MarginsEntity is responsible to give a valid entity.Margins.
func MarginsEntity() entity.Margins { _ = "STUB: not implemented"; return *new(entity.Margins) }

// ImageEntity is responsible to give a valid entity.Image.
func ImageEntity() entity.Image { _ = "STUB: not implemented"; return *new(entity.Image) }

// CellProp is responsible to give a valid props.Cell.
func CellProp() props.Cell { _ = "STUB: not implemented"; return *new(props.Cell) }

// Node is responsible to give a valid node.Node.
func Node(rootType string) *node.Node[core.Structure] { _ = "STUB: not implemented"; return nil }

// ColorProp is responsible to give a valid props.Color.
func ColorProp() props.Color { _ = "STUB: not implemented"; return *new(props.Color) }

// CheckboxProp is responsible to give a valid props.Checkbox.
func CheckboxProp() props.Checkbox { _ = "STUB: not implemented"; return *new(props.Checkbox) }

// LineProp is responsible to give a valid props.Line.
func LineProp() props.Line { _ = "STUB: not implemented"; return *new(props.Line) }

// SignatureProp is responsible to give a valid props.Signature.
func SignatureProp() props.Signature { _ = "STUB: not implemented"; return *new(props.Signature) }

// PageProp is responsible to give a valid props.PageNumber.
func PageProp() props.PageNumber { _ = "STUB: not implemented"; return *new(props.PageNumber) }
