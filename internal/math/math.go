package math

import (
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

type Math struct{}

// New create a Math.
func New() *Math {
	_ = "STUB: not implemented"

	// Resize adjusts the internal dimension of an element to occupy a percentage of the available space
	//   - inner: The inner dimensions of the element
	//   - outer: The outer dimensions of the element
	//   - percent: The percentage of the external dimension that can be occupied
	//   - justReferenceWidth: Indicates whether resizing should be done only in relation to width or in relation to width and height
	return nil
}

func (s *Math) Resize(inner *entity.Dimensions, outer *entity.Dimensions, percent float64, justReferenceWidth bool) *entity.Dimensions {
	_ = "STUB: not implemented"
	return nil
}

// GetInnerCenterCell define a inner cell formatted inside outer cell centered.
func (s *Math) GetInnerCenterCell(inner *entity.Dimensions, outer *entity.Dimensions) *entity.Cell {
	_ = "STUB: not implemented"
	return nil
}

// GetCenterCorrection return the correction of space in X or Y to
// centralize a line in relation with another line.
func (s *Math) GetCenterCorrection(outerSize, innerSize float64) float64 {
	_ = "STUB: not implemented"
	return 0
}
