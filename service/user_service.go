package service

import (
	"context"
	"golang-res-api-coba/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Delete(ctx context.Context, userId int) 
	FindById(ctx context.Context,  userId int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}