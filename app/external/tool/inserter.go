package tool

import "github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/restful/handlers"

type Inserter struct {
	productHandler handlers.ProductHandler
}

func NewInserter(productHandler handlers.ProductHandler) *Inserter {
	return &Inserter{
		productHandler: productHandler,
	}
}

func (i *Inserter) BulkInsertProducts()
