package gormrepos

import (
	"context"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/appctx"
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/gormer"
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/logger"

	"gorm.io/gorm"
)

// TxnDataSQL manages the transaction by implementing the Transaction Manager interface
type TxnDataSQL struct {
	db gormer.DBAdapter
}

// NewTxnDataSQL is the constructor function
func NewTxnDataSQL(db gormer.DBAdapter) *TxnDataSQL {
	return &TxnDataSQL{
		db: db,
	}
}

// TxnBegin begin a new transaction
func (tds *TxnDataSQL) TxnBegin(ctx context.Context) context.Context {
	txn := tds.db.Gormer().WithContext(ctx).Begin()
	logger.Debugf("TxnBegin - txn pointer %p", txn)
	return appctx.SetValue(ctx, appctx.TransactionContextKey, txn)
}

// TxnRollback rollback a transaction
func (tds *TxnDataSQL) TxnRollback(ctx context.Context) (err error) {
	return tds.GetTxn(ctx).(*gorm.DB).Rollback().Error
}

// TxnCommit commit a transaction
func (tds *TxnDataSQL) TxnCommit(ctx context.Context) (err error) {
	return tds.GetTxn(ctx).(*gorm.DB).Commit().Error
}

// GetTxn to get the current transaction of this service
func (tds *TxnDataSQL) GetTxn(ctx context.Context) interface{} {
	v := appctx.GetValue(ctx, appctx.TransactionContextKey)
	logger.Debugf("GetTxn - txn pointer %p", v)
	db, ok := v.(*gorm.DB)
	if !ok {
		panic("assign to *gorm.DB failed")
	}
	return db
}
