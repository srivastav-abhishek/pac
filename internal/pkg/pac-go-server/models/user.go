package models

type User struct {
	ID        string   `json:"id"`
	Username  string   `json:"username"`
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Email     string   `json:"email"`
	Groups    []string `json:"groups"`
}
