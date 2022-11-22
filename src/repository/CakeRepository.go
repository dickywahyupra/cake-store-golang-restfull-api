package repository

import (
	"cake-store-golang-restfull-api/helper"
	"cake-store-golang-restfull-api/src/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CakeRepository struct {
}

func NewCakeRepository() CakeRepositoryInterface {
	return &CakeRepository{}
}

func (repository *CakeRepository) Save(ctx context.Context, tx *sql.Tx, cake domain.Cake) domain.Cake {
	query := "INSERT INTO cakes (title, description, rating, image) VALUES (?,?,?,?)"

	result, err := tx.ExecContext(ctx, query, cake.Title, cake.Description, cake.Rating, cake.Image)
	helper.IfError(err)

	id, err := result.LastInsertId()
	helper.IfError(err)

	cake.Id = int(id)

	return cake
}

func (repository *CakeRepository) Update(ctx context.Context, tx *sql.Tx, cake domain.Cake) domain.Cake {
	query := "UPDATE cakes set title = ?, description = ?, rating = ?, image = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, cake.Title, cake.Description, cake.Rating, cake.Image, cake.Id)
	helper.IfError(err)

	return cake
}

func (repository *CakeRepository) Delete(ctx context.Context, tx *sql.Tx, cakeId int) {
	query := "DELETE FROM cakes WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, cakeId)
	helper.IfError(err)
}

func (repository *CakeRepository) FindById(ctx context.Context, tx *sql.Tx, cakeId int) (domain.Cake, error) {
	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?"

	rows, err := tx.QueryContext(ctx, query, cakeId)
	helper.IfError(err)

	defer rows.Close()

	cake := domain.Cake{}

	if rows.Next() {
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		helper.IfError(err)

		return cake, nil
	} else {
		return cake, errors.New("data not found")
	}
}

func (repository *CakeRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Cake {
	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes"

	rows, err := tx.QueryContext(ctx, query)
	helper.IfError(err)

	defer rows.Close()

	var cakes []domain.Cake

	for rows.Next() {
		cake := domain.Cake{}
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		helper.IfError(err)
		cakes = append(cakes, cake)
	}

	return cakes
}
