package fetcher

import (
	"context"
	"errors"
	"news-service/internal"

	"github.com/richarddes/newsapi-golang"
)

func (f *Fetcher) topHeadlines(ctx context.Context, lang string) (newsapi.TopHeadlinesResp, error) {
	if !internal.IsSupportedLang(lang) {
		return newsapi.TopHeadlinesResp{}, errors.New("Only specified lang isn't a valid lang")
	}

	// Parse the specified 2-letter ISO-639-1 code to a 2-letter ISO 3166-1 code.
	// Currently only en's parsed since that's the default language.
	if lang == "en" {
		lang = "us"
	}

	opts := newsapi.TopHeadlinesOpts{Country: lang, Category: "business"}
	r, err := f.Client.TopHeadlines(ctx, opts)
	if err != nil {
		return newsapi.TopHeadlinesResp{}, err
	}

	return r, nil
}
