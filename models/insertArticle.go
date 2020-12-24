package models

import (
	"context"
	"news-service/config"

	"github.com/lib/pq"
)

// InsertArticle adds a new article to the db.
func (db *DB) InsertArticle(ctx context.Context, art config.Article) error {
	stmt := "INSERT INTO news VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);"

	_, err := db.ExecContext(ctx, stmt, art.URL, art.PublisherID, art.PublisherName, art.Lang, art.Title, art.URLToImage, art.Author, art.Description, art.PublishedAt, art.Content)
	if err != nil {
		// We don't want to throw an error if the article already exists.
		if err, ok := err.(*pq.Error); ok {
			// Error code 23505 checks for a unique violation.
			if err.Code != "23505" {
				return err
			}
		}
	}

	return nil
}
