## Basic Usage
Full Example [example](example)

Full Example with Clean Architecture
[example](https://github.com/dreamph/go-clean-architecture-template)

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dreamph/dbre"
	"github.com/dreamph/dbre/adapters/bun"
	bunpg "github.com/dreamph/dbre/adapters/bun/connectors/pg"
	"github.com/dreamph/dbre/adapters/gorm"
	gormpg "github.com/dreamph/dbre/adapters/gorm/connectors/pg"
	"github.com/dreamph/dbre/example/domain"
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

	data := &domain.Country{
		Id:     "1",
		Code:   "C1",
		Name:   "Name",
		Status: 20,
	}

	// Create
	_, err = countryDbQuery.Create(ctx, data)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Update
	_, err = countryDbQuery.Update(ctx, data)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Find By PK
	_, err = countryDbQuery.FindByPK(ctx, &domain.Country{
		Id: "1",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Find One
	_, err = countryDbQuery.FindOne(ctx, &domain.Country{
		Code: "C1",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Delete
	err = countryDbQuery.Delete(ctx, &domain.Country{
		Id: "C1",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Upsert
	_, err = countryDbQuery.Upsert(ctx, data, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Find One Where
	wb := dbre.WhereBuilder{}
	wb.Where("code = ?", "C1")
	result, err := countryDbQuery.FindOneWhere(ctx, wb.WhereCauses())
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(result)

	// Query List Where
	wb = dbre.WhereBuilder{}
	wb.Where("status = ?", 20)

	list, total, err := countryDbQuery.QueryListWhere(ctx, wb.WhereCauses(), &dbre.Limit{Offset: 0, PageSize: 10}, []string{"name"})
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(list, total)

	// With Transaction
	err = dbTx.WithTx(ctx, func(ctx context.Context, appDB dbre.AppIDB) error {
		data2 := &domain.Country{
			Id:     "1",
			Code:   "C1",
			Name:   "Name",
			Status: 20,
		}
		_, err = countryDbQuery.WithTx(appDB).Create(ctx, data2)
		if err != nil {
			log.Fatalf(err.Error())
		}

		data2.Name = "Name2"
		data2.Status = 10

		_, err = countryDbQuery.WithTx(appDB).Update(ctx, data2)
		if err != nil {
			log.Fatalf(err.Error())
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
