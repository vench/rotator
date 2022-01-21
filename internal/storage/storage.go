package storage

import "github.com/vench/rotator/internal/entities"

// Feed persist data feed.
type Feed interface {
	// Upsert update feed item if exists or create.
	Upsert(item *entities.FeedItem) error
	// GetAll get list persists feed items.
	GetAll() ([]*entities.FeedItem, error)
}

// Block persist data block.
type Block interface {
	// Upsert update feed item if exists or create.
	Upsert(item *entities.BlockItem) error
	// GetAll get list persists feed items.
	GetAll() ([]*entities.BlockItem, error)
}
