package repository

import (
	"context"
	"database/sql"
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
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	user.Id = int(id)
	return user
}
func (reposiory *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	panic("err")
}
func (reposiory *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	panic("err")
}
func (reposiory *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	panic("err")
}
func (reposiory *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	panic("err")
}