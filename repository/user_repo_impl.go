package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-res-api-coba/helper"
	"golang-res-api-coba/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (reposiory *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO user(email, password, type) VALUES(?, ?, ?)"
	res, err := tx.ExecContext(ctx, SQL, user.Email, user.Password, user.Type) 
	helper.PanicError(err)
	id, err := res.LastInsertId()
	helper.PanicError(err)
	user.Id = int(id)
	return user
}
func (reposiory *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE user SET email = ?, type = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Email, user.Type, user.Id)
	helper.PanicError(err)
	return user
}
func (reposiory *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE FROM user WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicError(err)
}
func (reposiory *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT id, email, type FROM user where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
	 err :=rows.Scan(&user.Id, &user.Email, &user.Type)
		helper.PanicError(err)
	 return user, nil
	}else {
		return user, errors.New("User is Not Found")
	}
}
func (reposiory *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, email, type FROM user"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	defer rows.Close()

	var users  []domain.User
	for rows.Next() {
		user := domain.User{}
		err :=rows.Scan(&user.Id, &user.Email, &user.Type)
		helper.PanicError(err)
		users = append(users, user)
	}
	 return users
	
}