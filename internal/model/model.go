package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Todo struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Username  string `json:"username"`
}

type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
