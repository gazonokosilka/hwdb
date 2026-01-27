package httpserver

import (
	"context"
	"encoding/json"
	"net/http"

	"hw/feature_sql/connection"
	sqlrequestsgo "hw/feature_sql/sql_requests.go"
	"hw/models"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Нужен метод POST", http.StatusMethodNotAllowed)
		return
	}

	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Плохой запрос", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	conn, err := connection.ConnectionDB(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	if err = sqlrequestsgo.InsertRow(ctx, conn, emp.FullName, emp.Position); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Нужен метод GET", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()
	conn, err := connection.ConnectionDB(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	emp, err := sqlrequestsgo.SelectRow(ctx, conn)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	ctx := context.Background()
	conn, err := connection.ConnectionDB(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	err = sqlrequestsgo.DelRow(ctx, conn, id)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
