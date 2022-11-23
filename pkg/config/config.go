package config

import (
	"errors"
	"os"
)

type Config struct {
	APIURL string
	Token  string
}

const (
	DefaultAPIURL = "https://api.recontent.app/public"
)

func New() (*Config, error) {
	customAPIURL := os.Getenv("RECONTENT_API_URL")
	token := os.Getenv("RECONTENT_API_KEY")

	if len(token) == 0 {
		return nil, errors.New("RECONTENT_API_KEY is required")
	}

	usedAPIURL := DefaultAPIURL

	if len(customAPIURL) > 0 {
		usedAPIURL = customAPIURL
	}

	return &Config{
		APIURL: usedAPIURL,
		Token:  token,
	}, nil
}
