package gorm

import (
	"github.com/dreamph/dbre"
	"gorm.io/gorm"
)

type gormIDB struct {
	db *gorm.DB
}

func NewIDB(db *gorm.DB) dbre.AppIDB {
	return &gormIDB{db: db}
}

func (g *gormIDB) GetDB() any {
	return g.db
}

func (g *gormIDB) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
