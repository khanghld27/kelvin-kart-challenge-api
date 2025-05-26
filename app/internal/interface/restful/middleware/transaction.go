package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/appctx"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/transaction"
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/logger"
)

// TransactionMiddleware middleware to help manage the transaction
type TransactionMiddleware struct {
	manager transaction.Manager
}

// NewTransactionMiddleware constructor
func NewTransactionMiddleware(manager transaction.Manager) TransactionMiddleware {
	return TransactionMiddleware{
		manager: manager,
	}
}

// StartRequest start the transaction at the beginning of a request
func (mw *TransactionMiddleware) StartRequest(ctx *gin.Context) {
	newCtx := mw.manager.TxnBegin(ctx.Request.Context())
	ctx.Request = ctx.Request.WithContext(newCtx)
	ctx.Next()
}

// EndRequest get error to check if you need to commit or rollback
func (mw *TransactionMiddleware) EndRequest(ctx *gin.Context) {
	ctx.Next()
	err := appctx.GetValue(ctx.Request.Context(), appctx.ErrorContextKey)
	if p := recover(); p != nil {
		logger.Error("found p and rollback ", p)
		err := mw.manager.TxnRollback(ctx.Request.Context())
		if err != nil {
			return
		}
	} else if err != nil {
		logger.Debugf("found e and rollback %v", err)
		err := mw.manager.TxnRollback(ctx.Request.Context())
		if err != nil {
			return
		}
	} else {
		logger.Debugf("commit transaction %p", mw.manager.GetTxn(ctx.Request.Context()))
		err := mw.manager.TxnCommit(ctx.Request.Context())
		if err != nil {
			return
		}
	}
}

func (mw *TransactionMiddleware) StartToolRequest(ctx context.Context) context.Context {
	newCtx := mw.manager.TxnBegin(ctx)
	return newCtx
}

func (mw *TransactionMiddleware) EndToolRequest(ctx context.Context) {
	err := appctx.GetValue(ctx, appctx.ErrorContextKey)
	if p := recover(); p != nil {
		logger.Error("found p and rollback ", p)
		err := mw.manager.TxnRollback(ctx)
		if err != nil {
			return
		}
	} else if err != nil {
		logger.Debugf("found e and rollback %v", err)
		err := mw.manager.TxnRollback(ctx)
		if err != nil {
			return
		}
	} else {
		logger.Debugf("commit transaction %p", mw.manager.GetTxn(ctx))
		err := mw.manager.TxnCommit(ctx)
		if err != nil {
			return
		}
	}
}
