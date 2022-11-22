package repository

import (
	"cake-store-golang-restfull-api/src/model/domain"
	"context"
	"database/sql"
)

type CakeRepositoryInterface interface {
	Save(ctx context.Context, tx *sql.Tx, cake domain.Cake) domain.Cake
	Update(ctx context.Context, tx *sql.Tx, cake domain.Cake) domain.Cake
	Delete(ctx context.Context, tx *sql.Tx, cakeId int)
	FindById(ctx context.Context, tx *sql.Tx, cakeId int) (domain.Cake, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Cake
}
