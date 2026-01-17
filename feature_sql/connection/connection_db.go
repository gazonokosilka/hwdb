package connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnectionDB(ctx context.Context) (*pgx.Conn, error) {
	url := "postgres://postgres:admin@localhost:5432/hw"
	return pgx.Connect(ctx, url)
}
