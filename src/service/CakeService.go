package service

import (
	"cake-store-golang-restfull-api/helper"
	"cake-store-golang-restfull-api/src/exception"
	"cake-store-golang-restfull-api/src/model/domain"
	"cake-store-golang-restfull-api/src/model/request"
	"cake-store-golang-restfull-api/src/model/response"
	"cake-store-golang-restfull-api/src/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CakeService struct {
	CakeRepository repository.CakeRepositoryInterface
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewCakeService(cakeRepository repository.CakeRepositoryInterface, DB *sql.DB, validate *validator.Validate) CakeServiceInterface {
	return &CakeService{
		CakeRepository: cakeRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *CakeService) Create(ctx context.Context, request request.CakeCreateRequest) response.CakeCreateResponse {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	tx, err := service.DB.Begin()
	helper.IfError(err)

	defer helper.CommitOrRollback(tx)

	cake := domain.Cake{
		Title:       request.Title,
		Description: request.Description,
		Image:       request.Image,
		Rating:      request.Rating,
	}

	cake = service.CakeRepository.Save(ctx, tx, cake)

	return helper.ToCakeCreateResponse(cake)
}

func (service *CakeService) Update(ctx context.Context, request request.CakeUpdateRequest) response.CakeResponse {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	tx, err := service.DB.Begin()
	helper.IfError(err)

	defer helper.CommitOrRollback(tx)

	cake, err := service.CakeRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	cake.Title = request.Title
	cake.Description = request.Description
	cake.Rating = request.Rating
	cake.Image = request.Image

	cake = service.CakeRepository.Update(ctx, tx, cake)

	return helper.ToCakeResponse(cake)
}

func (service *CakeService) Delete(ctx context.Context, cakeId int) {
	tx, err := service.DB.Begin()
	helper.IfError(err)

	defer helper.CommitOrRollback(tx)

	_, err = service.CakeRepository.FindById(ctx, tx, cakeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CakeRepository.Delete(ctx, tx, cakeId)
}

func (service *CakeService) FindById(ctx context.Context, cakeId int) response.CakeResponse {
	tx, err := service.DB.Begin()
	helper.IfError(err)

	defer helper.CommitOrRollback(tx)

	cake, err := service.CakeRepository.FindById(ctx, tx, cakeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCakeResponse(cake)
}

func (service *CakeService) FindAll(ctx context.Context) []response.CakeResponse {
	tx, err := service.DB.Begin()
	helper.IfError(err)

	defer helper.CommitOrRollback(tx)

	cakes := service.CakeRepository.FindAll(ctx, tx)

	return helper.ToAllCakeResponse(cakes)
}
