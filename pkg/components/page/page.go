// Package page implements creation of pages.
package page

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Page struct {
	number int
	total  int
	rows   []core.Row
	config *entity.Config
	prop   props.PageNumber
}

// New is responsible to create a core.Page.
func New(ps ...props.PageNumber) core.Page { _ = "STUB: not implemented"; return *new(core.Page) }

// Render renders a Page into a PDF context.
func (p *Page) Render(provider core.Provider, cell entity.Cell) { _ = "STUB: not implemented"; return }

// SetConfig sets the Page configuration.
func (p *Page) SetConfig(config *entity.Config) { _ = "STUB: not implemented"; return }

// SetNumber sets the Page number and total.
func (p *Page) SetNumber(number int, total int) { _ = "STUB: not implemented"; return }

// GetNumber returns the Page number.
func (p *Page) GetNumber() int {
	_ = "STUB: not implemented"

	// Add adds one or more rows to the Page.
	return 0
}

func (p *Page) Add(rows ...core.Row) core.Page { _ = "STUB: not implemented"; return *new(core.Page) }

// GetRows returns the rows of the Page.
func (p *Page) GetRows() []core.Row {
	_ = "STUB: not implemented"

	// GetStructure returns the Structure of a Page.
	return nil
}

func (p *Page) GetStructure() *node.Node[core.Structure] { _ = "STUB: not implemented"; return nil }
