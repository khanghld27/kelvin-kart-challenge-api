package repositories

import (
	"context"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/models"
)

type ProductRepository interface {
	GetByID(ctx context.Context, id int) (*models.Product, error)
	GetProducts(ctx context.Context) ([]*models.Product, error)
	CreateProduct(ctx context.Context, product *models.Product) error
	BulkCreateProducts(ctx context.Context, products []*models.Product) error
}
