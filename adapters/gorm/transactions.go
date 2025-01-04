package gorm

import (
	"context"

	"github.com/dreamph/dbre/query"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type dbTx struct {
	db *gorm.DB
}

func NewDBTx(db *gorm.DB) query.DBTx {
	return &dbTx{db: db}
}

func (t *dbTx) WithTx(ctx context.Context, fn query.TxFn) (err error) {
	tx := t.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			if e, ok := r.(error); ok {
				err = errors.Wrap(e, "rollback tx")
			}
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	err = fn(ctx, NewIDB(tx))
	if err != nil {
		return err
	}
	return nil
}
