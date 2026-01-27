package models

type Employee struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	Position string `json:"position"`
}
