package web

type UserUpdateRequest struct {
	Id   int    `validate:"required"`
	Type string `validate:"required,max=200,min=1" json:"name"`
}