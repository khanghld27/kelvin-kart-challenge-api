package gormrepos

import (
	"context"

	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/appctx"
	"gorm.io/gorm"
)

type baseRepository struct{}

// DB to get the transaction to Database from context
func (r *baseRepository) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(appctx.TransactionContextKey)
	gormDB, ok := v.(*gorm.DB)
	if !ok {
		panic("can not get the gorm.DB in context")
	}
	return gormDB
}
