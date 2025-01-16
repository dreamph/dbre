package domain

import (
	"time"

	"github.com/guregu/null"
	"github.com/uptrace/bun"
)

var (
	_ = time.Second
	_ = null.Bool{}
)

/*
DB Table Details
-------------------------------------


Table: country
[ 0] id                                             VARCHAR              null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: -1      default: []
[ 1] code                                           VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
[ 2] name                                           VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
[ 3] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] description                                    VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
[ 5] other_field                                    VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "oBmMJPYcnsUkcIYOGuHLsRWRL",    "code": "BCqblMbsigdVuPGLUZCoWuKdt",    "name": "bMQtJJNaokqZvhqLxuOrAxWTZ",    "status": 64,    "description": "IowjpFMXLkJYNBpNPgGVAOTVx",    "otherField": "UlkHXLUgYWQrKIfUljaBTMXgw"}



*/

// Country struct is a row record of the country table in the public database
type Country struct {
	bun.BaseModel `bun:"table:country,alias:c" json:"-" swaggerignore:"true"`
	// [ 0] id                                             VARCHAR              null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Id string `bun:"id,pk" gorm:"primary_key;column:id;type:VARCHAR;" json:"id"`
	// [ 1] code                                           VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Code string `gorm:"column:code;type:VARCHAR;" json:"code"`
	// [ 2] name                                           VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Name string `gorm:"column:name;type:VARCHAR;" json:"name"`
	// [ 3] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Status int32 `gorm:"column:status;type:INT4;" json:"status"`
	// [ 4] description                                    VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Description null.String `gorm:"column:description;type:VARCHAR;" json:"description" swaggertype:"string"`
	// [ 5] other_field                                    VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	OtherField null.String `gorm:"column:other_field;type:VARCHAR;" json:"otherField" swaggertype:"string"`
}

// TableName sets the insert table name for this struct type
func (c *Country) TableName() string {
	return "country"
}
