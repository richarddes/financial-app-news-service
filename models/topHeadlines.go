package models

import (
	"context"
	"errors"
	"news-service/config"
	"news-service/internal"
)

// TopHeadlines returns the top headlines with the specified language in the db.
// It fails if the specified language isn't a supported language.
func (db *DB) TopHeadlines(ctx context.Context, lang string) ([]config.Article, error) {
	if !internal.IsSupportedLang(lang) {
		return nil, errors.New("The specified language isn't a supported language")
	}

	stmt := "SELECT publisherID,publisherName,author,description,publishedAt,title,url,urlToImage FROM news WHERE lang=$1;"

	rows, err := db.QueryContext(ctx, stmt, lang)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	arts := make([]config.Article, 0)
	for rows.Next() {
		th := new(config.Article)
		if err := rows.Scan(&th.PublisherID, &th.PublisherName, &th.Author, &th.Description, &th.PublishedAt, &th.Title, &th.URL, &th.URLToImage); err != nil {
			return nil, err
		}

		arts = append(arts, *th)
	}

	return arts, nil
}
