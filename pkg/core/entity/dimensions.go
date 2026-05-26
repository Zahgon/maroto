package entity

// Dimensions is the representation of a width and height.
type Dimensions struct {
	Width  float64
	Height float64
}

// AppendMap appends the dimensions to a map.
func (d *Dimensions) AppendMap(label string, m map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
