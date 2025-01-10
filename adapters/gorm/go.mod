module github.com/dreamph/dbre/adapters/gorm

go 1.23

replace github.com/dreamph/dbre => ../..

require (
	github.com/dreamph/dbre v0.0.0-20250110043151-cb4a6b9ab013
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pkg/errors v0.9.1
	go.uber.org/zap v1.27.0
	gorm.io/driver/postgres v1.5.11
	gorm.io/gorm v1.25.12
	moul.io/zapgorm2 v1.3.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.2 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
