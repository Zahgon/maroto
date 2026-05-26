package code

import (
	"errors"
	"image"

	libBarcode "github.com/boombuler/barcode"

	"github.com/johnfercher/maroto/v2/pkg/consts/barcode"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var (
	ErrCannotEncodePNG        = errors.New("cannot encode png")
	ErrCannotScaleBarcode     = errors.New("cannot scale barcode")
	ErrCannotEncodeQRcode     = errors.New("cannot encode qr code")
	ErrCannotEncodeDataMatrix = errors.New("cannot encode data matrix")
)

type Code struct{}

// New create a Code (Singleton).
func New() *Code {
	_ = "STUB: not implemented"

	// GenDataMatrix is responsible to generate a data matrix byte array.
	return nil
}

func (c *Code) GenDataMatrix(code string) (*entity.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenQr is responsible to generate a qr code byte array.
func (c *Code) GenQr(code string) (*entity.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenBar is responsible to generate a barcode byte array.
func (c *Code) GenBar(code string, _ *entity.Cell, prop *props.Barcode) (*entity.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getBarcodeClosure(
	barcodeType barcode.Type,
) func(code string) (libBarcode.BarcodeIntCS, error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Code) getImage(img image.Image) (*entity.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
