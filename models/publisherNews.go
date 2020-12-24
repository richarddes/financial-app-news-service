package models

import (
	"context"
	"errors"
	"news-service/config"
	"strings"
)

// PublisherNews fetches the news from every publisher with an id specified in publisherIDs
func (db *DB) PublisherNews(ctx context.Context, publisherIDs []string) ([]config.Article, error) {
	if len(publisherIDs) < 1 {
		return []config.Article{}, errors.New("No publisher IDs have been specified")
	}

	query := `SELECT publisherID,publisherName,author,description,publishedAt,title,url,urlToImage FROM news WHERE publisherID = ANY($1::TEXT[]);`

	param := "{" + strings.Join(publisherIDs, ",") + "}"

	rows, err := db.QueryContext(ctx, query, param)
	if err != nil {
		return []config.Article{}, err
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
