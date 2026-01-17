package sqlrequestsgo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DelRow(ctx context.Context, conn *pgx.Conn, books []int) error {
	sqlQuery := `
		DELETE FROM books
		WHERE id = ANY($1)
	`
	_, err := conn.Exec(ctx, sqlQuery, books)
	return err
}
