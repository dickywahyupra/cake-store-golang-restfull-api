package helper

import (
	"cake-store-golang-restfull-api/src/model/domain"
	"cake-store-golang-restfull-api/src/model/response"
)

func ToCakeResponse(cake domain.Cake) response.CakeResponse {
	return response.CakeResponse{
		Id:          cake.Id,
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
		CreatedAt:   cake.CreatedAt,
		UpdatedAt:   cake.UpdatedAt,
	}
}

func ToCakeCreateResponse(cake domain.Cake) response.CakeCreateResponse {
	return response.CakeCreateResponse{
		Id:          cake.Id,
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
	}
}

func ToAllCakeResponse(cakes []domain.Cake) []response.CakeResponse {
	var cakeResponses []response.CakeResponse

	for _, cake := range cakes {
		cakeResponses = append(cakeResponses, ToCakeResponse(cake))
	}

	return cakeResponses
}
