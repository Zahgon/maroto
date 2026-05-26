package entity

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
)

const trimSizeDefault = 10

// Image is the representation of an image that can be added to the pdf.
type Image struct {
	Bytes      []byte
	Extension  extension.Type
	Dimensions *Dimensions
}

// AppendMap adds the Image fields to the map.
func (i *Image) AppendMap(m map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }
