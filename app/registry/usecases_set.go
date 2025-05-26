package registry

import (
	"github.com/google/wire"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/usecases"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/usecases/interactor"
)

var (
	useCasesSet = wire.NewSet(
		interactor.NewProductUseCase,
		wire.Bind(new(usecases.ProductUseCase), new(*interactor.ProductUseCase)),
	)
)
