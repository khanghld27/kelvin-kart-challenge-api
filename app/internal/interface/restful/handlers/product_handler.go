package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/restful/presenter"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/usecases"
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/copier"
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
	var (
		productId int
		res       presenter.ProductInformation
		err       error
	)

	id := ctx.Param("id")
	// parse id to int
	productId, err = strconv.Atoi(id)
	if err != nil {
		h.SetError(ctx, err)
		return
	}

	resDto, err := h.productUseCase.GetProductByID(ctx.Request.Context(), productId)
	if err != nil {
		h.SetError(ctx, err)
		return
	}

	copier.MustCopy(&res, &resDto)
	h.SetData(ctx, res)
	h.SetMeta(ctx, presenter.MetaResponse{
		Code: http.StatusOK,
	})
}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	var (
		res []presenter.ProductInformation
		err error
	)

	resDto, err := h.productUseCase.GetProducts(ctx.Request.Context())
	if err != nil {
		h.SetError(ctx, err)
		return
	}

	copier.MustCopy(&res, &resDto)
	h.SetData(ctx, res)
	h.SetMeta(ctx, presenter.MetaResponse{
		Code: http.StatusOK,
	})
}
