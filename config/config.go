// Package config defines globally used interfaces and structs.
package config

import (
	"context"
	"time"
)

var (
	// SupportedLangs defines the supported news languages.
	// It should be set once the program starts
	SupportedLangs []string
	Publishers     map[string][]Publisher
)

type (
	// Env represents a collection of interfaces required for the handlers.
	Env struct {
		DB         Datastore
		NewsClient NewsFetcher
	}

	// Article represents an Article like it's saved in the datastore
	Article struct {
		PublisherName, URL, URLToImage, Title, Description string
		Author, PublishedAt, Content, Lang, PublisherID    string
	}

	Publisher struct {
		ID, Name, Description string
	}
)

type (
	// Datastore defines functions a datastore has to implement.
	Datastore interface {
		TopHeadlines(ctx context.Context, lang string) ([]Article, error)
		InsertArticle(ctx context.Context, art Article) error
		Clear(ctx context.Context) error
		EntriesExist(ctx context.Context) (bool, error)
		PublisherNews(ctx context.Context, publisherIDs []string) ([]Article, error)
	}

	// NewsFetcher defines functions a news fetcher has to implement.
	NewsFetcher interface {
		FetchAndSave(ctx context.Context, env *Env, interval time.Duration) error
		PublisherArticles(ctx context.Context, lang string, publisherIDs []string) ([]Article, error)
		Publishers(ctx context.Context, lang string) ([]Publisher, error)
	}
)
