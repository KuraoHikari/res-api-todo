package web

type CategoryCreateRequest struct {
	Email    string `validate:"required,min=1,max=200" json:"name"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}