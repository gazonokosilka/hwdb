package sqlrequestsgo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Pagination(ctx context.Context, conn *pgx.Conn, pageSize int, offset int) ([]BookModel, error) {
	sqlQuery := `
		SELECT * FROM books
		ORDER BY id ASC
		LIMIT $1 OFFSET $2 
	`
	rows, err := conn.Query(ctx, sqlQuery, pageSize, offset)
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

func GetTotalCount(ctx context.Context, conn *pgx.Conn) (int, error) {
	var count int
	query := `
		SELECT COUNT(*) FROM books
	`
	err := conn.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func ListPages(ctx context.Context, conn *pgx.Conn, N int) error {
	totalcount, err := GetTotalCount(ctx, conn)
	if err != nil {
		return err
	}
	totalPages := (totalcount + N - 1) / N

	pageNumber := 1
	offset := 0

	for pageNumber <= totalPages {
		books, err := Pagination(ctx, conn, pageNumber, offset)
		if err != nil {
			return err
		}
		fmt.Printf("Страница %d: ", pageNumber)
		if len(books) == 0 {
			fmt.Println("[]")
		} else {
			fmt.Println("[")
			for i, book := range books {
				fmt.Printf("%d %s", book.ID, book.Name)
				if i < len(books)-1 {
					fmt.Print(", ")
				}
			}
			fmt.Println("]")
		}
		offset += N
		pageNumber++
	}
	return nil
}
