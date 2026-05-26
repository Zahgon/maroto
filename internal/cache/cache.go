package cache

import (
	"errors"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

var (
	ErrCannotReadFile = errors.New("cannot read file")
	ErrImageNotFound  = errors.New("image not found")
)

// Cache is the interface to cache images.
type Cache interface {
	GetImage(value string, extension extension.Type) (*entity.Image, error)
	LoadImage(value string, extension extension.Type) error
	AddImage(value string, image *entity.Image)
}

type cache struct {
	images map[string]*entity.Image
	codes  map[string][]byte
}

// New is responsible to create a new Cache.
func New() Cache { _ = "STUB: not implemented"; return *new(Cache) }

// LoadImage loads an image from a file.
func (c *cache) LoadImage(file string, extension extension.Type) error {
	_ = "STUB: not implemented"
	return nil
}

// AddImage adds an image to the cache.
func (c *cache) AddImage(value string, image *entity.Image) { _ = "STUB: not implemented"; return }

// GetImage returns an image from the cache.
func (c *cache) GetImage(file string, extension extension.Type) (*entity.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
