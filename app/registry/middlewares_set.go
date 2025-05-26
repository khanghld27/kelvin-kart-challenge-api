package registry

import (
	"github.com/google/wire"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/persistence/rdbms/gormrepos"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/restful/middleware"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/transaction"
)

var (
	txnMwSet = wire.NewSet(
		singletonSet,
		gormrepos.NewTxnDataSQL,
		wire.Bind(new(transaction.Manager), new(*gormrepos.TxnDataSQL)),
		middleware.NewTransactionMiddleware,
	)
)
