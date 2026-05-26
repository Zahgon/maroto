package gofpdf

import (
	"errors"

	"github.com/johnfercher/maroto/v2/internal/cache"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var ErrCannotReadImageOptions = errors.New("could not read image options, maybe path/name is wrong")

type provider struct {
	fpdf       gofpdfwrapper.Fpdf
	font       core.Font
	text       core.Text
	code       core.Code
	image      core.Image
	line       core.Line
	checkbox   core.Checkbox
	cache      cache.Cache
	cellWriter cellwriter.CellWriter
	cfg        *entity.Config
}

// New is the constructor of provider for gofpdf
func New(dep *Dependencies) core.Provider { _ = "STUB: not implemented"; return *new(core.Provider) }

func (g *provider) AddText(text string, cell *entity.Cell, prop *props.Text) {
	_ = "STUB: not implemented"
	return
}

func (g *provider) GetLinesQuantity(text string, textProp *props.Text, colWidth float64) int {
	_ = "STUB: not implemented"
	return 0
}

func (g *provider) GetFontHeight(prop *props.Font) float64 { _ = "STUB: not implemented"; return 0 }

func (g *provider) AddLine(cell *entity.Cell, prop *props.Line) { _ = "STUB: not implemented"; return }

func (g *provider) AddCheckbox(label string, cell *entity.Cell, prop *props.Checkbox) {
	_ = "STUB: not implemented"
	return
}

func (g *provider) AddMatrixCode(code string, cell *entity.Cell, prop *props.Rect) {
	_ = "STUB: not implemented"
	return
}

func (g *provider) AddQrCode(code string, cell *entity.Cell, prop *props.Rect) {
	_ = "STUB: not implemented"
	return
}

func (g *provider) AddBarCode(code string, cell *entity.Cell, prop *props.Barcode) {
	_ = "STUB: not implemented"
	return
}

func (g *provider) AddImageFromFile(file string, cell *entity.Cell, prop *props.Rect) {
	_ = "STUB: not implemented"
	return
}

func (g *provider) AddImageFromBytes(bytes []byte, cell *entity.Cell, prop *props.Rect, extension extension.Type) {
	_ = "STUB: not implemented"
	return
}

func (g *provider) AddBackgroundImageFromBytes(bytes []byte, cell *entity.Cell, prop *props.Rect, extension extension.Type) {
	_ = "STUB: not implemented"
	return
}

func (g *provider) CreateRow(height float64) { _ = "STUB: not implemented"; return }

func (g *provider) SetProtection(protection *entity.Protection) { _ = "STUB: not implemented"; return }

func (g *provider) SetMetadata(metadata *entity.Metadata) { _ = "STUB: not implemented"; return }

// GetDimensionsByImage is responsible for obtaining the dimensions of an image
// If the image cannot be loaded, an error is returned
func (g *provider) GetDimensionsByImage(file string) (*entity.Dimensions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDimensionsByImageByte is responsible for obtaining the dimensions of an image
// If the image cannot be loaded, an error is returned
func (g *provider) GetDimensionsByImageByte(bytes []byte, extension extension.Type) (*entity.Dimensions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDimensionsByMatrixCode is responsible for obtaining the dimensions of an MatrixCode
// If the image cannot be loaded, an error is returned
func (g *provider) GetDimensionsByMatrixCode(code string) (*entity.Dimensions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDimensionsByQrCode is responsible for obtaining the dimensions of an QrCode
// If the image cannot be loaded, an error is returned
func (g *provider) GetDimensionsByQrCode(code string) (*entity.Dimensions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *provider) GenerateBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (g *provider) CreateCol(width, height float64, config *entity.Config, prop *props.Cell) {
	_ = "STUB: not implemented"
	return
}

func (g *provider) SetCompression(compression bool) { _ = "STUB: not implemented"; return }

func (g *provider) getBarcodeImageName(code string, prop *props.Barcode) string {
	_ = "STUB: not implemented"
	return ""
}

// loadImage is responsible for loading an codes
func (g *provider) loadCode(code, codeType string, generate func(code string) (*entity.Image, error)) (*entity.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loadImage is responsible for loading an image
func (g *provider) loadImage(file, extensionStr string) (*entity.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
