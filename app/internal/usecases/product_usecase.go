package usecases

import (
	"context"

	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/models"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/usecases/dto"
)

type ProductUseCase interface {
	GetProductByID(ctx context.Context, id int) (dto.ProductResponse, error)
	GetProducts(ctx context.Context) ([]dto.ProductResponse, error)
	CreateProduct(ctx context.Context, product *models.Product) error
}

type ProductImporterUseCase interface {
	ImportProducts(ctx context.Context) error
}
