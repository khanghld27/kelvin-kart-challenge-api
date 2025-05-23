package interactor

import (
	"context"

	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/repositories"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/usecases/dto"
)

type ProductUseCase struct {
	productRepo repositories.ProductRepository
}

func NewProductUseCase(productRepo repositories.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}

func (p *ProductUseCase) GetProductByID(ctx context.Context, id int) (dto.ProductResponse, error) {
	product, err := p.productRepo.GetByID(ctx, id)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.MapProductToResponse(product), nil
}

func (p *ProductUseCase) GetProducts(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := p.productRepo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	var productResponses []dto.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, dto.MapProductToResponse(product))
	}

	return productResponses, nil
}
