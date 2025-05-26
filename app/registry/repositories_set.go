package registry

import (
	"github.com/google/wire"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/repositories"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/persistence/rdbms/gormrepos"
)

var (
	repositoriesSet = wire.NewSet(
		gormrepos.NewProductRepository,
		wire.Bind(new(repositories.ProductRepository), new(*gormrepos.ProductRepository)),
	)
)
