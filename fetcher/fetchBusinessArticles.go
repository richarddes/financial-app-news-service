package fetcher

import (
	"context"
	"errors"
	"fmt"
	"news-service/config"

	"github.com/richarddes/newsapi-golang"
)

// PublisherArticles fetches the news from each publisher with an id specified in publisherIDs.
func (f *Fetcher) PublisherArticles(ctx context.Context, lang string, publisherIDs []string) ([]config.Article, error) {
	if len(publisherIDs) < 1 {
		return []config.Article{}, errors.New("No publisher IDs have been specified")
	}

	opts := newsapi.TopHeadlinesOpts{Sources: publisherIDs, PageSize: 100}

	arts, err := f.Client.TopHeadlines(ctx, opts)
	if err != nil {
		return []config.Article{}, err
	}

	as := make([]config.Article, len(arts.Articles))

	for i, art := range arts.Articles {
		fmtPubDate := fmt.Sprintf("%d.%d.%d", art.PublishedAt.Day(), art.PublishedAt.Month(), art.PublishedAt.Year())

		as[i] = config.Article{
			Author:        art.Author,
			Description:   art.Description,
			PublishedAt:   fmtPubDate,
			PublisherID:   art.Source.ID,
			PublisherName: art.Source.Name,
			Title:         art.Title,
			URL:           art.URL,
			URLToImage:    art.URLToImage,
			Content:       art.Content,
			Lang:          lang,
		}
	}

	return as, nil
}
