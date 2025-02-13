package pg

import (
	"context"
	"database/sql"

	"github.com/dreamph/dbre"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunotel"
	"go.uber.org/zap"

	"fmt"
)

type Options struct {
	Host           string
	Port           string
	DBName         string
	User           string
	Password       string
	ConnectTimeout int64
	Logger         *zap.Logger
	PoolOptions    *dbre.DbPoolOptions
	TraceEnable    bool
	Timezone       string
}

func Connect(options *Options) (*bun.DB, error) {
	connection := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		options.User,
		options.Password,
		options.Host,
		options.Port,
		options.DBName,
	)

	db, err := openSQLDB(false, connection, options.Timezone, options.PoolOptions)
	if err != nil {
		return nil, err
	}

	bunDB := bun.NewDB(db, pgdialect.New(), bun.WithDiscardUnknownColumns())
	if options.TraceEnable {
		bunDB.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName(options.DBName)))
	}

	bunDB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	/*	bunDB.AddQueryHook(bunzap.NewQueryHook(bunzap.QueryHookOptions{
		Logger: options.Logger,
	}))*/

	return bunDB, nil
}

func openSQLDB(usePgxPool bool, connection string, timezone string, poolOptions *dbre.DbPoolOptions) (*sql.DB, error) {
	if usePgxPool {
		return openDBFromPool(connection, timezone, poolOptions)
	}
	return openDB(connection, timezone, poolOptions)
}

func openDB(connection string, timezone string, poolOptions *dbre.DbPoolOptions) (*sql.DB, error) {
	config, err := pgx.ParseConfig(connection)
	if err != nil {
		return nil, err
	}

	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	config.RuntimeParams = setTimeZone(config.RuntimeParams, timezone)

	db := stdlib.OpenDB(*config)
	err = db.PingContext(context.Background())
	if err != nil {
		return nil, err
	}

	dbPool := dbre.DbPoolDefault
	if poolOptions != nil {
		dbPool = poolOptions
	}

	dbre.SetConnectionsPool(db, dbPool)

	return db, nil
}

func openDBFromPool(connection string, timezone string, poolOptions *dbre.DbPoolOptions) (*sql.DB, error) {
	poolConfig, err := pgxpool.ParseConfig(connection)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	poolConfig.ConnConfig.RuntimeParams = setTimeZone(poolConfig.ConnConfig.RuntimeParams, timezone)

	if poolOptions != nil {
		poolConfig.MinConns = int32(poolOptions.MaxIdleConns)
		poolConfig.MaxConns = int32(poolOptions.MaxOpenConns)
		poolConfig.MaxConnLifetime = poolOptions.ConnMaxLifetime
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDBFromPool(pool)
	err = db.PingContext(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setTimeZone(runtimeParams map[string]string, timezone string) map[string]string {
	if timezone != "" {
		runtimeParams["timezone"] = timezone
	} else {
		runtimeParams["timezone"] = "Asia/Bangkok"
	}
	return runtimeParams
}
