package feed

import (
	"fmt"
	"os"

	"github.com/vench/rotator/internal/entities"

	"github.com/vench/rotator/internal/config"

	jsoniter "github.com/json-iterator/go"
)

func New(conf *config.Feed) (*Service, error) {
	return &Service{
		conf: conf,
	}, nil
}

type Service struct {
	conf *config.Feed
}

func (s *Service) Load() error {
	// Todo read items
	data, err := os.ReadFile(s.conf.Path)
	if err != nil {
		return fmt.Errorf("failed to read file: %s, %v", s.conf.Path, err)
	}
	var list []*entities.FeedItem
	if err := jsoniter.Unmarshal(data, &list); err != nil {
		return fmt.Errorf("failed to unmarshal feed items: %v", err)
	}
	fmt.Println(list[0])

	// Todo read blocks

	return nil
}
