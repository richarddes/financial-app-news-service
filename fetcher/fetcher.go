package fetcher

import (
	"context"
	"errors"
	"fmt"
	"log"
	"news-service/config"
	"sync"
	"time"

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

// FetchAndSave first checks if news already exist in the db. If so, it waits for interval minutes. 
// Otherwise it fetches new data every interval and savaes the recvieved articles in the datastore
// specified in the env parameter.
// This function should be run in a seperate goroutine as it's an infinite loop and only terminates
// when the program terminates.
func (f *Fetcher) FetchAndSave(ctx context.Context, env *config.Env, interval time.Duration) error {
	exists, err := env.DB.EntriesExist(ctx)
	if err != nil {
		log.Panic(err)
	}

	if !exists {
		f.fetchAndSaveArts(ctx, env)
	}

	time.Sleep(interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	tc := ticker.C

	for {
		select {
		case <-tc:
			err := f.fetchAndSaveArts(ctx, env)
			if err != nil {
				log.Fatal(err)
			}

		case <-ctx.Done():
			return errors.New("Run out of time")
		}
	}
}

func (f *Fetcher) fetchAndSaveArts(ctx context.Context, env *config.Env) error {
	errc1 := f.fetchAndSaveTopHeadlines(ctx, env, "en")
	errc2 := f.fetchAndSaveTopHeadlines(ctx, env, "de")

	err := env.DB.Clear(ctx)
	if err != nil {
		return err
	}

	errc := fanInErrs(errc1, errc2)
	if err := <-errc; err != nil {
		return err
	}

	return nil
}

func (f *Fetcher) fetchAndSaveTopHeadlines(ctx context.Context, env *config.Env, lang string) <-chan error {
	var wg sync.WaitGroup
	errc := make(chan error, 1)

	wg.Add(1)

	go func() {
		resp, err := f.topHeadlines(ctx, lang)
		if err != nil {
			errc <- err
		}

		for _, art := range resp.Articles {
			fmtPubDate := fmt.Sprintf("%d.%d.%d", art.PublishedAt.Day(), art.PublishedAt.Month(), art.PublishedAt.Year())

			err := env.DB.InsertArticle(ctx, config.Article{
				Author:      art.Author,
				Description: art.Description,
				PublishedAt: fmtPubDate,
				Source:      art.Source.Name,
				Title:       art.Title,
				URL:         art.URL,
				URLToImage:  art.URLToImage,
				Content:     art.Content,
				Lang:        lang,
			})
			if err != nil {
				errc <- err
			}
		}

		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(errc)
	}()

	return errc
}

func fanInErrs(cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	wg.Add(len(cs))

	out := make(chan error)

	output := func(c <-chan error) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
