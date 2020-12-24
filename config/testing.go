package config

import (
	"context"
	"time"
)

type (
	mockNF struct{}

	mockDB struct {
		store map[string]string
	}
)

func (db *mockDB) TopHeadlines(ctx context.Context, lang string) ([]Article, error) {
	return []Article{}, nil
}

func (db *mockDB) InsertArticle(ctx context.Context, art Article) error {
	return nil
}

func (db *mockDB) Clear(ctx context.Context) error {
	return nil
}

func (db *mockDB) EntriesExist(ctx context.Context) (bool, error) {
	return true, nil
}

func (db *mockDB) PublisherNews(ctx context.Context, publisherIDs []string) ([]Article, error) {
	return []Article{}, nil
}

func (nf *mockNF) FetchAndSave(ctx context.Context, env *Env, interval time.Duration) error {
	return nil
}

func (nf *mockNF) PublisherArticles(ctx context.Context, lang string, publisherIDs []string) ([]Article, error) {
	return []Article{}, nil
}

func (nf *mockNF) Publishers(ctx context.Context, lang string) ([]Publisher, error) {
	return []Publisher{}, nil
}

// NewMockEnv returns a new Env with mock values instead of production values
func NewMockEnv() *Env {
	env := new(Env)

	db := new(mockDB)
	db.store = make(map[string]string)

	nf := new(mockNF)

	env.DB = db
	env.NewsClient = nf

	return env
}
