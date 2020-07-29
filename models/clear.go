package models

import "context"

// Clear deletes all news fromthe database,
func (db *DB) Clear(ctx context.Context) error {
	stmt := "TRUNCATE TABLE news;"

	_, err := db.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	return nil
}
