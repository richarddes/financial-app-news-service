package models

import (
	"context"
)

// EntriesExist checks if there're any news left in the db.
func (db *DB) EntriesExist(ctx context.Context) (bool, error) {
	var n int

	stmt := "SELECT Count(url) FROM news;"

	err := db.QueryRowContext(ctx, stmt).Scan(&n)
	if err != nil {
		return false, err
	}

	if n > 0 {
		return true, nil
	}

	return false, nil
}
