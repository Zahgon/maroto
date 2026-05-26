package cache

import (
	"sync"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

type mutexCache struct {
	inner      Cache
	imageMutex sync.RWMutex
}

// NewMutexDecorator is responsible to create a mutex decorator to read/write cache.
func NewMutexDecorator(cache Cache) Cache { _ = "STUB: not implemented"; return *new(Cache) }

// LoadImage adds a behavior to lock/unlock cache write.
func (c *mutexCache) LoadImage(file string, extension extension.Type) error {
	_ = "STUB: not implemented"
	return nil
}

// AddImage adds a behavior to lock/unlock cache write.
func (c *mutexCache) AddImage(value string, image *entity.Image) { _ = "STUB: not implemented"; return }

// GetImage adds a behavior to lock/unlock cache read.
func (c *mutexCache) GetImage(file string, extension extension.Type) (*entity.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
