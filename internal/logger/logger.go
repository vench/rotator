package logger

import (
	"fmt"
	"go.uber.org/zap"
)

func New(serviceName string) (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}
	return logger, nil
}
