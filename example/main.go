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
