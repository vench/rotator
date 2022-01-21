package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// App contains application settings.
type App struct {
	// Name application name.
	Name string

	// Feed contains feed settings.
	Feed Feed

	// HTTP todo name for rotator and etc
	HTTP HTTP
}

// Feed contains feed settings.
type Feed struct {
	// Path contains path to data place.
	Path string
}

// HTTP contains http server settings.
type HTTP struct {
	// Port server port.
	Port int
}

// New create App instance.
func New(configPath string) (*App, error) {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.SetConfigFile(configPath)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var appConfig App
	if err := v.Unmarshal(&appConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &appConfig, nil
}
