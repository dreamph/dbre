package bun

import (
	"github.com/dreamph/dbre/query"
	"github.com/uptrace/bun"
)

type bunIDB struct {
	db bun.IDB
}

func NewIDB(db bun.IDB) query.AppIDB {
	return &bunIDB{db: db}
}

func (b *bunIDB) GetDB() any {
	return b.db
}
