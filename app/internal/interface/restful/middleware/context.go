package middleware

import (
	"context"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/configs"

	"github.com/gin-gonic/gin"
)

// AddTimeout for context
func AddTimeout(ctx *gin.Context) {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctxWTimeout, cancel := context.WithTimeout(
		ctx.Request.Context(),
		configs.GetConfig().HTTPServer.Timeout,
	)
	ctx.Request = ctx.Request.WithContext(ctxWTimeout)
	defer cancel()
	ctx.Next()
}
