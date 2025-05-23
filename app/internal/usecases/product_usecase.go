package usecases

import (
	"context"

	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/usecases/dto"
)

type ProductUseCase interface {
	GetProductByID(ctx context.Context, id int) (dto.ProductResponse, error)
	GetProducts(ctx context.Context) ([]dto.ProductResponse, error)
}
