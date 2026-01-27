package sqlrequestsgo

import (
	"context"
	"hw/models"

	"github.com/jackc/pgx/v5"
)

func SelectRow(con context.Context, conn *pgx.Conn) ([]models.Employee, error) {
	sqlQuery := `
		SELECT * FROM employee
		ORDER BY id ASC
	`
	rows, err := conn.Query(con, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employee := make([]models.Employee, 0)

	for rows.Next() {
		var emp models.Employee

		err := rows.Scan(
			&emp.Id,
			&emp.FullName,
			&emp.Position,
		)
		if err != nil {
			return nil, err
		}

		employee = append(employee, emp)
	}
	return employee, nil

}
