package gorm

import (
	"github.com/dreamph/dbre/query"
	"gorm.io/gorm"
)

type gormIDB struct {
	db *gorm.DB
}

func NewIDB(db *gorm.DB) query.AppIDB {
	return &gormIDB{db: db}
}

func (g *gormIDB) GetDB() any {
	return g.db
}
