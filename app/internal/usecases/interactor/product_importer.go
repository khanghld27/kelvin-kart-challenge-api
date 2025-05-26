package interactor

import (
	"context"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/gateway"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/models"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/repositories"
)

type ProductImporterUseCase struct {
	ExternalAPI       gateway.ExternalProductAPI
	ProductRepository repositories.ProductRepository
}

func NewProductImporterUseCase(api gateway.ExternalProductAPI, repo repositories.ProductRepository) *ProductImporterUseCase {
	return &ProductImporterUseCase{
		ExternalAPI:       api,
		ProductRepository: repo,
	}
}

func (uc *ProductImporterUseCase) ImportProducts(ctx context.Context) error {
	scrapedProducts, err := uc.ExternalAPI.FetchProducts(ctx)
	if err != nil {
		return err
	}

	products := make([]*models.Product, 0, len(scrapedProducts))
	for _, sp := range scrapedProducts {
		product := &models.Product{
			Id:       sp.Id,
			Name:     sp.Name,
			Category: sp.Category,
			Price:    sp.Price,
		}
		if sp.Image != nil {
			product.Image = &models.Image{
				Thumbnail: sp.Image.Thumbnail,
				Mobile:    sp.Image.Mobile,
				Tablet:    sp.Image.Tablet,
				Desktop:   sp.Image.Desktop,
			}
		}

		products = append(products, product)
	}

	if err = uc.ProductRepository.BulkCreateProducts(ctx, products); err != nil {
		return err
	}

	return nil
}
