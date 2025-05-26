package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/restful/presenter"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/usecases"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	BaseHandler
	productUseCase usecases.ProductUseCase
}

func NewProductHandler(productUseCase usecases.ProductUseCase) ProductHandler {
	return ProductHandler{
		productUseCase: productUseCase,
	}
}

func (h *ProductHandler) GetProductByID(ctx *gin.Context) {
	productID := ctx.Param("id")
	// parse productID to int or appropriate type if needed
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		h.SetError(ctx, err)
		return
	}

	product, err := h.productUseCase.GetProductByID(ctx.Request.Context(), productIDInt)
	if err != nil {
		h.SetError(ctx, err)
		return
	}
	h.SetData(ctx, product)
	h.SetMeta(ctx, presenter.MetaResponse{
		Code: http.StatusOK,
	})
}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := h.productUseCase.GetProducts(ctx.Request.Context())
	if err != nil {
		h.SetError(ctx, err)
		return
	}

	h.SetData(ctx, products)
	h.SetMeta(ctx, presenter.MetaResponse{
		Code: http.StatusOK,
	})
}

func (h *ProductHandler) CreateProduct(ctx context.Context) {
	return
}
