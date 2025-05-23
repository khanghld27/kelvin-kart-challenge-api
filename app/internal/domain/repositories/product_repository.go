package repositories

import (
	"context"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/models"
)

type ProductRepository interface {
	GetByID(ctx context.Context, id int) (*models.Product, error)
	GetProducts(ctx context.Context) ([]*models.Product, error)
}
