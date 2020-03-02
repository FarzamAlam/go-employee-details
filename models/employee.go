package models

type Employee struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Department   string `json:"department"`
}
