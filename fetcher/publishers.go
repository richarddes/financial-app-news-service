package fetcher

import (
	"context"
	"news-service/config"

	"github.com/richarddes/newsapi-golang"
)

func (f *Fetcher) Publishers(ctx context.Context, lang string) ([]config.Publisher, error) {
	opts := newsapi.SourcesOpts{Category: "business", Language: lang}

	resp, err := f.Client.Sources(ctx, opts)
	if err != nil {
		return []config.Publisher{}, err
	}

	sourceIds := make([]config.Publisher, len(resp.Sources))

	for i, source := range resp.Sources {
		sourceIds[i] = config.Publisher{ID: source.ID, Name: source.Name, Description: source.Description}
	}

	return sourceIds, nil
}
