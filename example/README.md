## Basic Usage
Full Example with Clean Architecture
[example](https://github.com/dreamph/go-clean-architecture-template)

Domain generate by [smallnest/gen](https://github.com/smallnest/gen)
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
	Upsert(ctx context.Context, obj *domain.Country, specifyUpdateFields []string) (*domain.Country, error)
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

func (r *countryRepository) Upsert(ctx context.Context, obj *domain.Country, specifyUpdateFields []string) (*domain.Country, error) {
	return r.query.Upsert(ctx, obj, specifyUpdateFields)
}

func (r *countryRepository) UpdateForce(ctx context.Context, obj *domain.Country) (*domain.Country, error) {
	return r.query.UpdateForce(ctx, obj)
}

func (r *countryRepository) Delete(ctx context.Context, id string) error {
	return r.query.Delete(ctx, &domain.Country{Id: id})
}

func (r *countryRepository) FindByID(ctx context.Context, id string) (*domain.Country, error) {
	return r.query.FindByPK(ctx, &domain.Country{Id: id})
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
				"code":   "code",
				"name":   "name",
				"status": "status",
			},
			Sort: obj.Sort,
			DefaultSort: &dbre.Sort{
				SortBy:        "name",
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
	"github.com/dreamph/dbre/example/core/models"
	"github.com/dreamph/dbre/example/domain/repomodels"

	"github.com/dreamph/dbre/example/domain"
	"github.com/dreamph/dbre/example/repository"
	"go.uber.org/zap"
)

func getBunDB(logger *zap.Logger) (dbre.AppIDB, dbre.DBTx, error) {
	bunDB, err := bunpg.Connect(&bunpg.Options{
		Host:           "127.0.0.1",
		Port:           "5432",
		DBName:         "dream",
		User:           "dream",
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
		DBName:         "dream",
		User:           "dream",
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
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer logger.Sync()

	appDB, dbTx, err := getBunDB(logger)
	//appDB, dbTx, err := getGormDB(logger)
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
		Id:     "1",
		Code:   "C1",
		Name:   "Name",
		Status: 20,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Simple Usage with Repository
	countryRepository := repository.NewCountryRepository(appDB)

	_, err = countryRepository.Create(ctx, &domain.Country{
		Id:     "12",
		Code:   "C12",
		Name:   "Name",
		Status: 20,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, _, err = countryRepository.List(ctx, &repomodels.CountryListCriteria{
		Limit: &models.PageLimit{
			PageNumber: 1,
			PageSize:   20,
		},
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	// With Transaction
	err = dbTx.WithTx(ctx, func(ctx context.Context, appDB dbre.AppIDB) error {
		_, err = countryRepository.WithTx(appDB).Create(ctx, &domain.Country{
			Id:     "13",
			Code:   "C13",
			Name:   "Name",
			Status: 20,
		})
		if err != nil {
			return err
		}

		_, err = countryRepository.WithTx(appDB).Create(ctx, &domain.Country{
			Id:     "21",
			Code:   "C31",
			Name:   "Name",
			Status: 20,
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


Buy Me a Coffee
=======
[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/dreamph)