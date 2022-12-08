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
	defer func(){
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

	return  web.UserResponse{
		Id: user.Id,
		Email: user.Email,
		Type: user.Type,
	}
}
func(service *UserServiceImpl)Update(ctx context.Context, request web.UserCreateRequest) web.UserResponse{
	panic("err")
}

func(service *UserServiceImpl)Delete(ctx context.Context, userId int) {
	panic("err")
}

func(service *UserServiceImpl)FindById(ctx context.Context,  userId int) web.UserResponse {
	tx, err :=service.DB.Begin()
	helper.PanicError(err)

	defer func(){
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

	user, err := service.UserRepository.FindById(ctx, tx , userId)
	helper.PanicError(err)

	return web.UserResponse{
		Id: user.Id,
		Email: user.Email,
		Type: user.Type,
	}
}
func(service *UserServiceImpl)FindAll(ctx context.Context) []web.UserResponse {
	panic("err")
}