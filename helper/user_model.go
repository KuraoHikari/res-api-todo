package helper

import (
	"golang-res-api-coba/model/domain"
	"golang-res-api-coba/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:   user.Id,
		Email: user.Email,
		Type: user.Type,
	}
}

// func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
// 	var CategoryResponse []web.CategoryResponse

// 	for _, category := range categories{
// 		CategoryResponse = append(CategoryResponse, ToCategoryResponse(category))
// 	}
// 	return CategoryResponse
// }