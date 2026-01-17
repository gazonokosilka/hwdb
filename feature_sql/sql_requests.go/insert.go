package sqlrequestsgo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, book BookModel) error {
	sqlQuery := `
		INSERT INTO books (name, author, resolution, reales_year, completed, time_on_shelf)
		VALUES($1, $2, $3, $4, $5, $6)
	`
	_, err := conn.Exec(ctx,
		sqlQuery,
		book.Name,
		book.Author,
		book.Resolution,
		book.Reales_year,
		book.Completed,
		book.Time_on_shelf)
	if err != nil {
		return err
	}
	return nil
}
