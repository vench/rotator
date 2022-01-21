package entities

import "encoding/json"

// FeedItem contain information about advert content.
type FeedItem struct {
	ID           uint64
	Coefficients []float64
	Assets       json.RawMessage
}
