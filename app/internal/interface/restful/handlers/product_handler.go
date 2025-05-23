package handlers

import "github.com/khanghld27/kelvin-kart-challenge-api/app/internal/usecases"

type ProductHandler struct {
	BaseHandler
	productUseCase usecases.ProductUseCase
}
