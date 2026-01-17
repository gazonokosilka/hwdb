package sqlrequestsgo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRow(con context.Context, conn *pgx.Conn) ([]BookModel, error) {
	sqlQuery := `
		SELECT * FROM books
		ORDER BY id ASC
	`
	rows, err := conn.Query(con, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]BookModel, 0)

	for rows.Next() {
		var book BookModel

		err := rows.Scan(
			&book.ID,
			&book.Name,
			&book.Author,
			&book.Resolution,
			&book.Reales_year,
			&book.Completed,
			&book.Time_on_shelf,
			&book.Time_completed,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}
	return books, nil

}
