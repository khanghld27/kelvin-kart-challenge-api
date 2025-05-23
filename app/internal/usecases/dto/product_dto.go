package dto

import "github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/models"

type ProductResponse struct {
	Id       string         `json:"id"`
	Image    *ImageResponse `json:"image"`
	Name     string         `json:"name"`
	Category string         `json:"category"`
	Price    float64        `json:"price"`
}

type ImageResponse struct {
	Thumbnail string `json:"thumbnail"`
	Mobile    string `json:"mobile"`
	Tablet    string `json:"tablet"`
	Desktop   string `json:"desktop"`
}

func MapProductToResponse(product *models.Product) ProductResponse {
	response := ProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Category: product.Category,
		Price:    product.Price,
	}

	if product.Image != nil {
		response.Image = &ImageResponse{
			Thumbnail: product.Image.Thumbnail,
			Mobile:    product.Image.Mobile,
			Tablet:    product.Image.Tablet,
			Desktop:   product.Image.Desktop,
		}
	}

	return response
}
