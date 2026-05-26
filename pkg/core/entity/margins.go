package entity

// Margins is the representation of a margin.
type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

// AppendMap appends the margins to a map.
func (m *Margins) AppendMap(mp map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
