package registry

import (
	"github.com/google/wire"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/restful/handlers"
)

var (
	productHandlerSet = wire.NewSet(
		singletonSet,
		repositoriesSet,
		useCasesSet,
		handlers.NewProductHandler,
	)
)
