package gofpdf

import (
	"errors"

	"github.com/google/uuid"
	"github.com/phpdave11/gofpdf"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var ErrCouldNotRegisterImageOptions = errors.New("could not register image options, maybe path/name is wrong")

type Image struct {
	pdf  gofpdfwrapper.Fpdf
	math core.Math
}

// NewImage create an Image.
func NewImage(pdf gofpdfwrapper.Fpdf, math core.Math) *Image { _ = "STUB: not implemented"; return nil }

// GetImageInfo is responsible for loading the image in PDF and returning its information.
func (s *Image) GetImageInfo(img *entity.Image, extension extension.Type) (*gofpdf.ImageInfoType, uuid.UUID) {
	_ = "STUB: not implemented"
	return nil, *new(uuid.UUID)
}

// Add use a byte array to add image to PDF.
func (s *Image) Add(img *entity.Image, cell *entity.Cell, margins *entity.Margins,
	prop *props.Rect, extension extension.Type, flow bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Image) addImageToPdf(imageLabel string, info *gofpdf.ImageInfoType, cell *entity.Cell, margins *entity.Margins,
	prop *props.Rect, flow bool,
) {
	_ = "STUB: not implemented"
	return
}
