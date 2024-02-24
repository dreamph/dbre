package main

import (
	"context"
	"log"

	"github.com/dreamph/dbre/example/domain"
	"github.com/dreamph/dbre/example/repository"
	"github.com/dreamph/dbre/query"
	"github.com/dreamph/dbre/query/bun"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf(err.Error())
	}

	bunDB, err := bun.Connect(&bun.Options{
		Host:           "127.0.0.1",
		Port:           "5432",
		DBName:         "DB1",
		User:           "user1",
		Password:       "password",
		ConnectTimeout: 2000,
		Logger:         logger,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer bun.Close(bunDB)

	appDB := &query.AppIDB{BunDB: bunDB}
	dbTx := bun.NewDBTx(bunDB)

	ctx := context.Background()
	countryRepository := repository.NewCountryRepository(appDB)

	//Simple Usage
	_, err = countryRepository.Create(ctx, &domain.Country{
		ID:     "1",
		NameEn: "",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	// With Transaction
	err = dbTx.WithTx(ctx, func(ctx context.Context, appDB *query.AppIDB) error {
		_, err = countryRepository.WithTx(appDB).Create(ctx, &domain.Country{
			ID:     "1",
			NameEn: "",
		})
		_, err = countryRepository.WithTx(appDB).Create(ctx, &domain.Country{
			ID:     "2",
			NameEn: "",
		})
		return nil
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
}
