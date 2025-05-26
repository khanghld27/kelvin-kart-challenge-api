//go:build wireinject

package registry

import (
	"github.com/google/wire"
	domain_gateway "github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/gateway"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/repositories"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/persistence/rdbms/gormrepos"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/scrapers/gateway"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/usecases/interactor"
)

func InitializeProductImporter(apiBaseURL string) *interactor.ProductImporterUseCase {
	wire.Build(
		gormrepos.NewProductRepository,
		wire.Bind(new(repositories.ProductRepository), new(*gormrepos.ProductRepository)),
		wire.Bind(new(domain_gateway.ExternalProductAPI), new(*gateway.HTTPExternalProductAPI)),
		gateway.NewHTTPExternalProductAPI,
		interactor.NewProductImporterUseCase,
	)

	return nil
}
