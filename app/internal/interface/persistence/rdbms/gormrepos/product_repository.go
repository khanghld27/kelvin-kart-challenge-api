package gormrepos

import (
	"context"

	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/models"
)

type ProductRepository struct {
	baseRepository
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	product := &models.Product{}
	// Preload image
	err := r.DB(ctx).Take(product, id).Preload("Images").Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) GetProducts(ctx context.Context) ([]*models.Product, error) {
	products := []*models.Product{}
	err := r.DB(ctx).Find(&products).Preload("Images").Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
