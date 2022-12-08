package web

type UserResponse struct {
	Id    int    `json:"id"`
	Email string `json:"name"`
	Type  string `json:"type"`
}