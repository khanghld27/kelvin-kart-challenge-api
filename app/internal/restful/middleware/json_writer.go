package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/appctx"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/restful/presenter"
	"net/http"
)

// JSONWriterMiddleware follow jsonapi.org
// Respond check errors, prepare meta and respond data
func JSONWriterMiddleware(ctx *gin.Context) {
	ctx.Next()

	// Check error if exists
	// Base on error/success to return meta object
	var (
		res      presenter.Response
		httpCode int
	)
	appErr := appctx.GetValue(ctx.Request.Context(), appctx.ErrorContextKey)
	if appErr != nil {
		_processAppError(&res, appErr)
		httpCode = res.Meta.Code
	}

	// Respond the data object/array
	data := appctx.GetValue(ctx.Request.Context(), appctx.DataContextKey)
	if data != nil {
		res.Data = data
	}
	meta := appctx.GetValue(ctx.Request.Context(), appctx.MetaContextKey)
	if meta != nil {
		metaRes, ok1 := meta.(presenter.MetaResponse)
		if ok1 {
			res.Meta = metaRes
			httpCode = metaRes.Code
		}
	}
	if res.IsEmpty() {
		ctx.JSON(http.StatusNoContent, nil)
	} else {
		ctx.JSON(httpCode, res)
	}
}

func _processAppError(res *presenter.Response, appErr interface{}) {
	bindingErr := _catchBindingError(appErr.(error))
	if bindingErr != nil {
		errors.As(bindingErr, &res.Errors)
		res.Meta = presenter.MetaResponse{
			Code:    http.StatusBadRequest,
			Message: "error when binding the request",
		}
	}
}

func _catchBindingError(appErr error) error {
	var errs presenter.ErrorResponses
	var invalidValidationError *validator.InvalidValidationError
	if errors.As(appErr, &invalidValidationError) {
		errs.Append(presenter.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Detail: "invalid validation error",
		})
		return errs
	}

	var vldrErr validator.ValidationErrors
	if errors.As(appErr, &vldrErr) {
		errs.FromValidationErrors(vldrErr)
		return errs
	}

	return nil
}
