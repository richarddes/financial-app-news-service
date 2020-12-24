package fetcher

import (
	"errors"
	"news-service/config"

	"github.com/richarddes/newsapi-golang"
)

type Fetcher struct {
	config.NewsFetcher
	Client newsapi.Client
}

// New returns a new Fetcher with the apiKey as the NewsAPI key.
func New(apiKey string) (*Fetcher, error) {
	if apiKey == "" {
		return nil, errors.New("The apiKey must have a value")
	}

	c := newsapi.Client{APIKey: apiKey}

	f := new(Fetcher)
	f.Client = c

	return f, nil
}
