// Package fontrepository implements font repository.
package fontrepository

import (
	"errors"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

var ErrCannotReadFile = errors.New("cannot read file")

// Repository is the abstraction to load custom fonts.
type Repository interface {
	AddUTF8Font(family string, style fontstyle.Type, file string) Repository
	AddUTF8FontFromBytes(family string, style fontstyle.Type, bytes []byte) Repository
	Load() ([]entity.CustomFont, error)
}

// FontRepository manages custom fonts to be loaded into a document.
type FontRepository struct {
	customFonts []*customFont
}

// New creates a new Repository.
func New() Repository {
	_ = "STUB: not implemented"
	return *

	// AddUTF8Font adds a custom font to the repository.
	new(Repository)
}

func (r *FontRepository) AddUTF8Font(family string, style fontstyle.Type, file string) Repository {
	_ = "STUB: not implemented"
	return *new(Repository)
}

// AddUTF8FontFromBytes adds a custom font to the repository from a byte slice.
func (r *FontRepository) AddUTF8FontFromBytes(family string, style fontstyle.Type, bytes []byte) Repository {
	_ = "STUB: not implemented"
	return *new(Repository)
}

// Load loads all custom fonts, reading file contents from disk where needed.
func (r *FontRepository) Load() ([]entity.CustomFont, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
