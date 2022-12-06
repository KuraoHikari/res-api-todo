package repository

import (
	"context"
	"database/sql"

	"github.com/KuraoHikari/golang-res-api/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx sql.Tx, user domain.User)
	FindById(ctx context.Context, tx sql.Tx, userId int) domain.User
	FindAll(ctx context.Context, tx sql.Tx)
}