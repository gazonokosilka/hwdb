package main

import (
	"fmt"
	httpserver "hw/http_server"
	"net/http"
)

func main() {
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			httpserver.CreateEmployee(w, r)
		case http.MethodGet:
			httpserver.GetAllEmployee(w, r)
		case http.MethodDelete:
			httpserver.DeleteEmployee(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("запуск сервера")
	if err := http.ListenAndServe(":5050", nil); err != nil {
		panic(err)
	}
}
