package transaction

import (
	"context"
)

type Manager interface {
	TxnBegin(ctx context.Context) context.Context
	TxnCommit(ctx context.Context) error
	TxnRollback(ctx context.Context) error
	GetTxn(ctx context.Context) interface{}
}
