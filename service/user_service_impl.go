package service

import (
	"context"
	"database/sql"
	"golang-res-api-coba/helper"
	"golang-res-api-coba/model/domain"
	"golang-res-api-coba/model/web"
	"golang-res-api-coba/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB				*sql.DB
	Validate		*validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validate.Validate) *UserServiceImpl{
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func(service *UserServiceImpl)Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer func(){ //defer helper.CommitRollback(tx)
		err := recover()
		if err != nil {
			errRollback := tx.Rollback()
			helper.PanicError(errRollback)
			panic(err)
		} else {
			errCommit := tx.Commit()
			helper.PanicError(errCommit)
		}
	}()
	user := domain.User{
		Email : request.Email,
		Password: request.Password,
		Type: request.Type,
	}
	user = service.UserRepository.Save(ctx, tx, user)

	return  web.UserResponse{ //return helper.ToUserResponse(user)
		Id: user.Id,
		Email: user.Email,
		Type: user.Type,
	}
}
func(service *UserServiceImpl)Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse{
	tx, err :=service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(tx)
	user, err := service.UserRepository.FindById(ctx, tx , request.Id)
	helper.PanicError(err)
	user.Type = request.Type
	user = service.UserRepository.Update(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func(service *UserServiceImpl)Delete(ctx context.Context, userId int) {
	tx, err :=service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(tx)
	user, err := service.UserRepository.FindById(ctx, tx , userId)
	helper.PanicError(err)
	service.UserRepository.Delete(ctx, tx, user)
}

func(service *UserServiceImpl)FindById(ctx context.Context,  userId int) web.UserResponse {
	tx, err :=service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(tx)
	user, err := service.UserRepository.FindById(ctx, tx , userId)
	helper.PanicError(err)
	return helper.ToUserResponse(user)
}
func(service *UserServiceImpl)FindAll(ctx context.Context) []web.UserResponse {
	tx, err :=service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(tx)
	users := service.UserRepository.FindAll(ctx, tx )
	var UserResponse []web.UserResponse

	for _, user := range users{
		UserResponse = append(UserResponse, helper.ToUserResponse(user))
	}
	return UserResponse
}