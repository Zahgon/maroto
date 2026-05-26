// Package signature implements creation of signatures.
package signature

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Signature struct {
	value  string
	prop   props.Signature
	config *entity.Config
}

// New is responsible to create an instance of a Signature.
func New(value string, ps ...props.Signature) core.Component {
	_ = "STUB: not implemented"
	return *new(core.Component)
}

// NewCol is responsible to create an instance of a Signature wrapped in a Col.
func NewCol(size int, value string, ps ...props.Signature) core.Col {
	_ = "STUB: not implemented"
	return *new(core.Col)
}

// NewRow is responsible to create an instance of a Signature wrapped in a Row.
func NewRow(height float64, value string, ps ...props.Signature) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// NewRow is responsible to create an instance of a Signature wrapped in a automatic Row.
func NewAutoRow(value string, ps ...props.Signature) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// Render renders a Signature into a PDF context.
func (s *Signature) Render(provider core.Provider, cell *entity.Cell) {
	_ = "STUB: not implemented"
	return
}

// GetStructure returns the Structure of a Signature.
func (s *Signature) GetStructure() *node.Node[core.Structure] {
	_ = "STUB: not implemented"
	return nil
}

// GetHeight returns the height that the signature will have in the PDF
func (s *Signature) GetHeight(provider core.Provider, _ *entity.Cell) float64 {
	_ = "STUB: not implemented"
	return 0
}

// SetConfig sets the config.
func (s *Signature) SetConfig(config *entity.Config) { _ = "STUB: not implemented"; return }
