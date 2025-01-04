package bun

import (
	"context"

	"github.com/dreamph/dbre/query"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type dbTx struct {
	db *bun.DB
}

func NewDBTx(db *bun.DB) query.DBTx {
	return &dbTx{db: db}
}

func (t *dbTx) WithTx(ctx context.Context, fn query.TxFn) (err error) {
	tx, err := t.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			if e, ok := r.(error); ok {
				err = errors.Wrap(e, "rollback tx")
			}
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(ctx, NewIDB(tx))
	if err != nil {
		return err
	}
	return nil
}
