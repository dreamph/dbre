package bun

import (
	"github.com/dreamph/dbre"
	"github.com/uptrace/bun"
)

type bunIDB struct {
	db bun.IDB
}

func NewIDB(db bun.IDB) dbre.AppIDB {
	return &bunIDB{db: db}
}

func (b *bunIDB) GetDB() any {
	return b.db
}

func (b *bunIDB) Close() error {
	bunDB, ok := b.db.(*bun.DB)
	if !ok {
		err := bunDB.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
