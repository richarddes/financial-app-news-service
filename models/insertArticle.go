package models

import (
	"context"
	"news-service/config"
)

// InsertArticle adds a new article to the db.
func (db *DB) InsertArticle(ctx context.Context, art config.Article) error {
	stmt := "INSERT INTO news VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9);"

	_, err := db.ExecContext(ctx, stmt, art.URL, art.Lang, art.Title, art.Source, art.URLToImage, art.Author, art.Description, art.PublishedAt, art.Content)
	if err != nil {
		return err
	}

	return nil
}
