//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package registry

import (
	"github.com/google/wire"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/restful/handlers"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/restful/middleware"
)

func TransactionMiddleware() middleware.TransactionMiddleware {
	wire.Build(txnMwSet)
	return middleware.TransactionMiddleware{}
}

func ProductHandler() handlers.ProductHandler {
	wire.Build(productHandlerSet)
	return handlers.ProductHandler{}
}
