package fixture

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

type TestFont struct {
	Family string
	Style  fontstyle.Type
	File   string
	Bytes  []byte
}

func (t TestFont) GetFamily() string { _ = "STUB: not implemented"; return "" }

func (t TestFont) GetStyle() fontstyle.Type { _ = "STUB: not implemented"; return *new(fontstyle.Type) }

func (t TestFont) GetFile() string { _ = "STUB: not implemented"; return "" }

func (t TestFont) GetBytes() []byte { _ = "STUB: not implemented"; return nil }
