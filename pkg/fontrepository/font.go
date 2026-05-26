package fontrepository

import "github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

type customFont struct {
	family string
	style  fontstyle.Type
	file   string
	bytes  []byte
}

func (c *customFont) GetFamily() string { _ = "STUB: not implemented"; return "" }
func (c *customFont) GetStyle() fontstyle.Type {
	_ = "STUB: not implemented"
	return *new(fontstyle.Type)
}
func (c *customFont) GetFile() string  { _ = "STUB: not implemented"; return "" }
func (c *customFont) GetBytes() []byte { _ = "STUB: not implemented"; return nil }
