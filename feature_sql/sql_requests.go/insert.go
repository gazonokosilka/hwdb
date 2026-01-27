package sqlrequestsgo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, fullName, position string) error {
	sqlQuery := `
		INSERT INTO employee (full_name, position)
		VALUES($1, $2)
	`
	_, err := conn.Exec(ctx,
		sqlQuery,
		fullName,
		position)
	if err != nil {
		return err
	}
	return nil
}
