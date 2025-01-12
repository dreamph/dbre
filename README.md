## Basic Usage
Full Example [example](example)

Full Example with Clean Architecture
[example](https://github.com/dreamph/go-clean-architecture-template/blob/main/internal/modules/company/usecase/company_example_db_transaction_usecase.go)
# Domain gen by https://github.com/smallnest/gen
```go
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
```

# Repo models
```go
package repomodels

type CountryListCriteria struct {
    Status int32             `json:"status" example:"20"`
    Limit  *models.PageLimit `json:"limit"`
    Sort   *models.Sort      `json:"sort"`
}
```

# Repository
```go
package repository

import (
	"context"

	"github.com/dreamph/dbre"
	"github.com/dreamph/dbre/adapters/bun"
	"github.com/dreamph/dbre/example/core/utils"
	"github.com/dreamph/dbre/example/domain"
	"github.com/dreamph/dbre/example/domain/repomodels"
)

type CountryRepository interface {
	WithTx(db dbre.AppIDB) CountryRepository

	Create(ctx context.Context, obj *domain.Country) (*domain.Country, error)
	Update(ctx context.Context, obj *domain.Country) (*domain.Country, error)
	UpdateForce(ctx context.Context, obj *domain.Country) (*domain.Country, error)
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*domain.Country, error)
	FindOne(ctx context.Context, obj *domain.Country) (*domain.Country, error)

	List(ctx context.Context, obj *repomodels.CountryListCriteria) (*[]domain.Country, int64, error)
}

type countryRepository struct {
	query dbre.DB[domain.Country]
}

func NewCountryRepository(db dbre.AppIDB) CountryRepository {
	return &countryRepository{
		query: bun.New[domain.Country](db),
	}
}

func (r *countryRepository) WithTx(tx dbre.AppIDB) CountryRepository {
	return NewCountryRepository(tx)
}

func (r *countryRepository) Create(ctx context.Context, obj *domain.Country) (*domain.Country, error) {
	return r.query.Create(ctx, obj)
}

func (r *countryRepository) Update(ctx context.Context, obj *domain.Country) (*domain.Country, error) {
	return r.query.Update(ctx, obj)
}

func (r *countryRepository) UpdateForce(ctx context.Context, obj *domain.Country) (*domain.Country, error) {
	return r.query.UpdateForce(ctx, obj)
}

func (r *countryRepository) Delete(ctx context.Context, id string) error {
	return r.query.Delete(ctx, &domain.Country{ID: id})
}

func (r *countryRepository) FindByID(ctx context.Context, id string) (*domain.Country, error) {
	return r.query.FindByPK(ctx, &domain.Country{ID: id})
}

func (r *countryRepository) FindOne(ctx context.Context, obj *domain.Country) (*domain.Country, error) {
	return r.query.FindOne(ctx, obj)
}

func (r *countryRepository) List(ctx context.Context, obj *repomodels.CountryListCriteria) (*[]domain.Country, int64, error) {
	result := &[]domain.Country{}
	whereBuilder := dbre.NewWhereBuilder()

	if obj.Status != 0 {
		whereBuilder.Where("status = ?", obj.Status)
	}

	whereCauses := whereBuilder.WhereCauses()
	total, err := r.query.CountWhere(ctx, whereCauses)
	if err != nil {
		return nil, 0, err
	}
	if total > 0 {
		sortSQL, err := dbre.SortSQL(&dbre.SortParam{
			SortFieldMapping: map[string]string{
				"id":     "id",
				"nameEn": "name_en",
				"nameTh": "name_th",
				"code":   "code",
				"status": "status",
			},
			Sort: obj.Sort,
			DefaultSort: &dbre.Sort{
				SortBy:        "nameEn",
				SortDirection: dbre.DESC,
			},
		})
		if err != nil {
			return nil, 0, err
		}

		result, err = r.query.ListWhere(ctx, whereCauses, utils.ToQueryLimit(obj.Limit), []string{sortSQL})
		if err != nil {
			return nil, 0, err
		}
	}

	return result, total, nil
}



```


# Main
```go
package main

import (
	"context"
	"log"

	"github.com/dreamph/dbre"
	"github.com/dreamph/dbre/adapters/bun"
	bunpg "github.com/dreamph/dbre/adapters/bun/connectors/pg"
	"github.com/dreamph/dbre/adapters/gorm"
	gormpg "github.com/dreamph/dbre/adapters/gorm/connectors/pg"

	"github.com/dreamph/dbre/example/domain"
	"github.com/dreamph/dbre/example/repository"
	"go.uber.org/zap"
)

func getBunDB(logger *zap.Logger) (dbre.AppIDB, dbre.DBTx, error) {
	bunDB, err := bunpg.Connect(&bunpg.Options{
		Host:           "127.0.0.1",
		Port:           "5432",
		DBName:         "DB1",
		User:           "user1",
		Password:       "password",
		ConnectTimeout: 2000,
		Logger:         logger,
	})
	if err != nil {
		return nil, nil, err
	}

	appDB := bun.NewIDB(bunDB)
	dbTx := bun.NewDBTx(bunDB)

	return appDB, dbTx, nil
}

func getGormDB(logger *zap.Logger) (dbre.AppIDB, dbre.DBTx, error) {
	bunDB, err := gormpg.Connect(&gormpg.Options{
		Host:           "127.0.0.1",
		Port:           "5432",
		DBName:         "DB1",
		User:           "user1",
		Password:       "password",
		ConnectTimeout: 2000,
		Logger:         logger,
	})
	if err != nil {
		return nil, nil, err
	}

	appDB := gorm.NewIDB(bunDB)
	dbTx := gorm.NewDBTx(bunDB)

	return appDB, dbTx, nil
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf(err.Error())
	}

	appDB, dbTx, err := getBunDB(logger) // or getGormDB(logger) for GORM
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer func() {
		err := appDB.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}()

	ctx := context.Background()

	//Simple Usage
	countryDbQuery := bun.New[domain.Country](appDB)
	_, err = countryDbQuery.Create(ctx, &domain.Country{
		ID:     "1",
		NameEn: "NameEn1",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Simple Usage with Repository
	countryRepository := repository.NewCountryRepository(appDB)

	_, err = countryRepository.Create(ctx, &domain.Country{
		ID:     "1",
		NameEn: "NameEn1",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	// With Transaction
	err = dbTx.WithTx(ctx, func(ctx context.Context, appDB dbre.AppIDB) error {
		_, err = countryRepository.WithTx(appDB).Create(ctx, &domain.Country{
			ID:     "1",
			NameEn: "NameEn1",
		})
		if err != nil {
			return err
		}

		_, err = countryRepository.WithTx(appDB).Create(ctx, &domain.Country{
			ID:     "2",
			NameEn: "NameEn2",
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
}



```
