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
	"github.com/google/uuid"

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

	//Simple Usage directly without Repository
	countryDbQuery := bun.New[domain.Country](appDB)
	_, err = countryDbQuery.Create(ctx, &domain.Country{
		Id:     uuid.New().String(),
		Code:   "C1",
		Name:   "Name",
		Status: 20,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Simple Usage with Repository
	countryRepository := repository.NewCountryRepository(appDB)

	id := uuid.New().String()
	_, err = countryRepository.Create(ctx, &domain.Country{
		Id:     id,
		Code:   "C12",
		Name:   "Name",
		Status: 20,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	data, err := countryRepository.FindByID(ctx, id)
	if err != nil {
		log.Fatalf(err.Error())
	}
	data.Name = "Name Updated"
	_, err = countryRepository.Update(ctx, data)
	if err != nil {
		log.Fatalf(err.Error())
	}

	data.Name = "Name Updated2"
	_, err = countryRepository.Upsert(ctx, data, []string{"name"})
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
			Id:     uuid.New().String(),
			Code:   "C13",
			Name:   "Name",
			Status: 20,
		})
		if err != nil {
			return err
		}

		_, err = countryRepository.WithTx(appDB).Create(ctx, &domain.Country{
			Id:     uuid.New().String(),
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
