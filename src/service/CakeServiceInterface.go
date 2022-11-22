package service

import (
	"cake-store-golang-restfull-api/src/model/request"
	"cake-store-golang-restfull-api/src/model/response"
	"context"
)

type CakeServiceInterface interface {
	Create(ctx context.Context, request request.CakeCreateRequest) response.CakeCreateResponse
	Update(ctx context.Context, request request.CakeUpdateRequest) response.CakeResponse
	Delete(ctx context.Context, cakeId int)
	FindById(ctx context.Context, cakeId int) response.CakeResponse
	FindAll(ctx context.Context) []response.CakeResponse
}
