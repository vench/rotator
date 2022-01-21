package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"time"

	"go.uber.org/zap"

	"github.com/vench/rotator/internal/logger"

	"github.com/vench/rotator/internal/server/http"

	"github.com/chapsuk/grace"

	"github.com/vench/rotator/internal/config"
	"github.com/vench/rotator/internal/service/feed"
	"golang.org/x/sync/errgroup"
)

// TODO this application read data source and return advert content by http endpoint
func main() {

	rand.Seed(time.Now().UnixNano())

	configPath := flag.String("config", "config/example.yaml", "path ro config file")
	flag.Parse()

	appConfig, err := config.New(*configPath)
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	ll, err := logger.New(appConfig.Name)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer ll.Sync()

	ctx := grace.ShutdownContext(context.Background())

	feed, err := feed.New(&appConfig.Feed)
	if err != nil {
		log.Fatalf("failed to create feed service: %v", err)
	}

	if err := feed.Load(); err != nil {
		log.Fatalf("failed to load feed service: %v", err)
	}

	s, err := http.NewServer(ll, appConfig)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	gr, appctx := errgroup.WithContext(ctx)
	gr.Go(func() error {
		return s.Serve(appctx)
	})

	if err := gr.Wait(); err != nil {
		ll.Error("failed to wait", zap.Error(err))
	}

	ll.Info("application has been shutdown")
}
