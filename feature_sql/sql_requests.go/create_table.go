package sqlrequestsgo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {

	sqlQuery := `
		CREATE TABLE IF NOT EXISTS books(
			id SERIAL PRIMARY KEY,
			name VARCHAR(200) NOT NULL,
			author VARCHAR(200) NOT NULL,
			resolution VARCHAR(200),
			reales_year INTEGER NOT NULL,
			completed BOOLEAN,
			time_on_shelf TIMESTAMP NOT NULL,
			time_comleted TIMESTAMP,

			UNIQUE(name)
		)
	`
	_, err := conn.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}
