package sqlrequestsgo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn, book BookModel) error {
	sqlQuery := `
		UPDATE books
		SET name = $1, author = $2, resolution = $3, reales_year = $4, completed = $5, time_on_shelf = $6, time_comleted = $7
		WHERE id = $8
	`
	_, err := conn.Exec(ctx, sqlQuery,
		book.Name,
		book.Author,
		book.Resolution,
		book.Reales_year,
		book.Completed,
		book.Time_on_shelf,
		book.Time_completed,
		book.ID,
	)
	return err
}
