package models

type UpdateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"Message"`
	Results any    `json:"Results"`
}
type Users struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}
