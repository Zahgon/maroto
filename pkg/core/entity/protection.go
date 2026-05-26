package entity

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/protection"
)

// Protection is the representation of a pdf protection.
type Protection struct {
	Type          protection.Type
	UserPassword  string
	OwnerPassword string
}

// AppendMap adds the Protection fields to the map.
func (p *Protection) AppendMap(m map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
