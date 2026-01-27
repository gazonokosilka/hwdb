package sqlrequestsgo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DelRow(ctx context.Context, conn *pgx.Conn, emp string) error {
	sqlQuery := `
		DELETE FROM employee
		WHERE id = $1
	`
	_, err := conn.Exec(ctx, sqlQuery, emp)
	return err
}
