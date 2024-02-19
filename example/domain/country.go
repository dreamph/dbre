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
[ 0] id                                             VARCHAR(45)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 45      default: []
[ 1] phone_code                                     VARCHAR(20)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
[ 2] code                                           VARCHAR(20)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
[ 3] name_th                                        VARCHAR(100)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
[ 4] name_en                                        VARCHAR(100)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
[ 5] iso_2                                          VARCHAR(5)           null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 5       default: []
[ 6] iso_3                                          VARCHAR(5)           null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 5       default: []
[ 7] iso_numeric_3                                  VARCHAR(5)           null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 5       default: []
[ 8] nationality_code                               VARCHAR(5)           null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 5       default: []
[ 9] nationality_name_en                            VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
[10] nationality_name_short_en                      VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
[11] name_short_en                                  VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
[12] reporting_currency                             VARCHAR(20)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
[13] currency_symbol                                VARCHAR(45)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 45      default: []
[14] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[15] seq                                            INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "gHUMRjhhrNEeCFVThKqcSBUXv",    "phoneCode": "emcFVvJpgQZfIGCUghxUlfKsR",    "code": "PcduWkbujpWAqfpjYkSiYovVC",    "nameTh": "HwVoSykLmqhLNwFFLtGqPMUmJ",    "nameEn": "VJaPFyjbVVtmdcxrMvHHDypEy",    "iso2": "rrMNOyIQyShVlRinbDsjTLwuw",    "iso3": "gVjaglpAuYoDCWhpmAUlsVKlc",    "isoNumeric3": "qWLvJJKojsMHNRuWrcsiGiMib",    "nationalityCode": "QxAIOuVAYCmIYJWuSpgYZXghE",    "nationalityNameEn": "VUSmCdErIToOvOpFooMNAAogr",    "nationalityNameShortEn": "dlaKAfNPSWSKEjOKkhTMhlNBD",    "nameShortEn": "QPNkFWoMJyIujWaklVucwuYOs",    "reportingCurrency": "OmYRlfGPhWuqnAnDOZITUXPXj",    "currencySymbol": "npwKEaaniobUKRYTZJjtZQAns",    "status": 68,    "seq": 30}



*/

// Country struct is a row record of the country table in the sodaudev database
type Country struct {
	bun.BaseModel `bun:"table:country,alias:c" json:"-" swaggerignore:"true"`
	//[ 0] id                                             VARCHAR(45)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 45      default: []
	ID string `bun:"id,pk" gorm:"primary_key;column:id;type:VARCHAR;size:45;" json:"id"`
	//[ 1] mobile_country_code                                     VARCHAR(20)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
	MobileCountryCode string `gorm:"column:mobile_country_code;type:VARCHAR;size:20;" json:"mobileCountryCode"`
	//[ 2] code                                           VARCHAR(20)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
	Code string `gorm:"column:code;type:VARCHAR;size:20;" json:"code"`
	//[ 3] name_th                                        VARCHAR(100)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
	NameTh string `gorm:"column:name_th;type:VARCHAR;size:100;" json:"nameTh"`
	//[ 4] name_en                                        VARCHAR(100)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
	NameEn string `gorm:"column:name_en;type:VARCHAR;size:100;" json:"nameEn"`
	//[ 5] iso2_code                                         VARCHAR(5)           null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 5       default: []
	Iso2Code null.String `gorm:"column:iso2_code;type:VARCHAR;size:5;" json:"iso2"`
	//[ 6] iso3_code                                          VARCHAR(5)           null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 5       default: []
	Iso3Code null.String `gorm:"column:iso3_code;type:VARCHAR;size:5;" json:"iso3"`
	//[ 7] iso_numeric3_code                                  VARCHAR(5)           null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 5       default: []
	IsoNumeric3Code null.String ` gorm:"column:iso_numeric3_code;type:VARCHAR;size:5;" json:"isoNumeric3"`
	//[ 8] nationality_code                               VARCHAR(5)           null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 5       default: []
	NationalityCode null.String `gorm:"column:nationality_code;type:VARCHAR;size:5;" json:"nationalityCode"`
	//[ 9] nationality_name_en                            VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
	NationalityNameEn null.String `gorm:"column:nationality_name_en;type:VARCHAR;size:100;" json:"nationalityNameEn"`
	//[10] nationality_name_short_en                      VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
	NationalityNameShortEn null.String `gorm:"column:nationality_name_short_en;type:VARCHAR;size:100;" json:"nationalityNameShortEn"`
	//[11] name_short_en                                  VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
	NameShortEn null.String `gorm:"column:name_short_en;type:VARCHAR;size:100;" json:"nameShortEn"`
	//[12] reporting_currency                             VARCHAR(20)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
	ReportingCurrency null.String `gorm:"column:reporting_currency;type:VARCHAR;size:20;" json:"reportingCurrency"`
	//[13] currency_symbol                                VARCHAR(45)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 45      default: []
	CurrencySymbol null.String `gorm:"column:currency_symbol;type:VARCHAR;size:45;" json:"currencySymbol"`
	//[14] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Status int32 `gorm:"column:status;type:INT4;" json:"status"`
	//[15] seq                                            INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Seq null.Int `gorm:"column:seq;type:INT4;" json:"seq"`
}

// TableName sets the insert table name for this struct type
func (c *Country) TableName() string {
	return "country"
}
