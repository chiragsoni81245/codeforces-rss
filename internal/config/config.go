package config

import (
	"os"
	"strconv"
	"strings"
)

// Config struct holds user preferences
type Config struct {
	Tags      []string
	MinRating int
	MaxRating int
}

// LoadConfig loads the configuration from environment variables or `config.json`
func LoadConfig() (*Config) {
	var cfg Config

	// Try to read from environment variables
	if tags, exists := os.LookupEnv("CF_TAGS"); exists {
		cfg.Tags = strings.Split(tags, ",")
	} else {
		cfg.Tags = []string{"dp", "graphs"} // Default
	}

	if minRating, exists := os.LookupEnv("CF_MIN_RATING"); exists {
		cfg.MinRating, _ = strconv.Atoi(minRating)
	} else {
		cfg.MinRating = 800 // Default
	}

	if maxRating, exists := os.LookupEnv("CF_MAX_RATING"); exists {
		cfg.MaxRating, _ = strconv.Atoi(maxRating)
	} else {
		cfg.MaxRating = 1500 // Default
	}

	return &cfg
}

