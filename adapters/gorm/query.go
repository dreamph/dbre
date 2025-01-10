package gorm

import (
	"context"
	"strings"

	"github.com/dreamph/dbre"
	"github.com/dreamph/dbre/adapters/gorm/utils"
	"gorm.io/gorm"
)

type dbQuery[T any] struct {
	DB *gorm.DB
}

func New[T any](db dbre.AppIDB) dbre.DB[T] {
	return &dbQuery[T]{DB: dbre.GetDB[*gorm.DB](db)}
}

func (q *dbQuery[T]) WithTx(tx dbre.AppIDB) dbre.DB[T] {
	return New[T](tx)
}

func (q *dbQuery[T]) Count(ctx context.Context, whereObj *T) (int64, error) {
	var obj T
	db := WithContext(ctx, q.DB)
	db = db.Model(&obj).Where(whereObj)

	var total int64
	db = db.Count(&total)
	if err := db.Error; err != nil {
		return 0, utils.DbError(err)
	}
	return total, nil
}

func (q *dbQuery[T]) List(ctx context.Context, whereObj *T) (*[]T, error) {
	var result []T
	db := WithContext(ctx, q.DB)
	db = db.Model(&result).Where(whereObj)

	db = db.Find(&result)
	if err := db.Error; err != nil {
		return nil, utils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) CountWhere(ctx context.Context, whereCauses *[]dbre.WhereCause) (int64, error) {
	var obj T
	db := WithContext(ctx, q.DB)
	db = db.Model(&obj)

	addWhere(db, whereCauses)

	var total int64
	db = db.Count(&total)
	if err := db.Error; err != nil {
		return 0, utils.DbError(err)
	}
	return total, nil
}

func (q *dbQuery[T]) ListWhere(ctx context.Context, whereCauses *[]dbre.WhereCause, limit *dbre.Limit, sortBy []string) (*[]T, error) {
	var result []T
	db := WithContext(ctx, q.DB)

	addWhere(db, whereCauses)

	if limit != nil {
		db.Limit(int(limit.PageSize)).Offset(int(limit.Offset))
	}
	if sortBy != nil {
		db.Order(strings.Join(sortBy, ","))
	}

	db = db.Find(&result)
	if err := db.Error; err != nil {
		return nil, utils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) QueryListWhere(ctx context.Context, whereCauses *[]dbre.WhereCause, limit *dbre.Limit, sortBy []string) (*[]T, int64, error) {
	var result []T
	db := WithContext(ctx, q.DB)

	addWhere(db, whereCauses)

	var total int64
	db = db.Count(&total)
	if err := db.Error; err != nil {
		return nil, 0, utils.DbError(err)
	}

	if limit != nil {
		db.Limit(int(limit.PageSize)).Offset(int(limit.Offset))
	}
	if sortBy != nil {
		db.Order(strings.Join(sortBy, ","))
	}

	db = db.Find(&result)
	if err := db.Error; err != nil {
		return nil, 0, utils.DbError(err)
	}
	return &result, total, nil
}

func (q *dbQuery[T]) RawQuery(ctx context.Context, sqlQuery string, params []interface{}, result interface{}) error {
	db := WithContext(ctx, q.DB)
	db.Raw(sqlQuery, params...).Scan(result)
	if err := db.Error; err != nil {
		return utils.DbError(err)
	}
	return nil
}

func (q *dbQuery[T]) RawExec(ctx context.Context, sqlQuery string, params []interface{}) (int64, error) {
	db := WithContext(ctx, q.DB)
	dbe := db.Exec(sqlQuery, params...)
	if err := dbe.Error; err != nil {
		return 0, utils.DbError(err)
	}
	return dbe.RowsAffected, nil
}

func (q *dbQuery[T]) Create(ctx context.Context, obj *T) (*T, error) {
	db := WithContext(ctx, q.DB)
	db = db.Create(&obj)
	if err := db.Error; err != nil {
		return nil, utils.DbError(err)
	}
	return obj, nil
}

func (q *dbQuery[T]) CreateList(ctx context.Context, obj *[]T) (*[]T, error) {
	db := WithContext(ctx, q.DB)
	db = db.Create(&obj)
	if err := db.Error; err != nil {
		return nil, utils.DbError(err)
	}
	return obj, nil
}

func (q *dbQuery[T]) Update(ctx context.Context, obj *T) (*T, error) {
	db := WithContext(ctx, q.DB)
	db = db.Updates(obj)
	if err := db.Error; err != nil {
		return nil, utils.DbError(err)
	}
	return obj, nil
}

func (q *dbQuery[T]) UpdateList(ctx context.Context, obj *[]T) (*[]T, error) {
	for _, row := range *obj {
		o := row
		_, err := q.Update(ctx, &o)
		if err != nil {
			return nil, utils.DbError(err)
		}
	}
	return obj, nil
}

func (q *dbQuery[T]) UpdateForce(ctx context.Context, obj *T) (*T, error) {
	db := WithContext(ctx, q.DB)
	db = db.Select("*").Updates(obj)
	if err := db.Error; err != nil {
		return nil, utils.DbError(err)
	}
	return obj, nil
}

func (q *dbQuery[T]) FindByPK(ctx context.Context, obj *T) (*T, error) {
	var result T
	db := WithContext(ctx, q.DB)
	if err := db.Where(obj).First(&result).Error; err != nil {
		return nil, utils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) FirstWhere(ctx context.Context, whereCauses *[]dbre.WhereCause, sortBy []string) (*T, error) {
	var result T
	db := WithContext(ctx, q.DB)

	addWhere(db, whereCauses)

	if sortBy != nil {
		db.Order(strings.Join(sortBy, ","))
	}

	db = db.First(&result)
	if err := db.Error; err != nil {
		return nil, utils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) FindOneWhere(ctx context.Context, whereCauses *[]dbre.WhereCause) (*T, error) {
	var result T
	db := WithContext(ctx, q.DB)

	addWhere(db, whereCauses)

	db = db.Find(&result)
	if err := db.Error; err != nil {
		return nil, utils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) FindOne(ctx context.Context, obj *T) (*T, error) {
	var result T
	db := WithContext(ctx, q.DB)
	if err := db.Where(obj).First(&result).Error; err != nil {
		return nil, utils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) Delete(ctx context.Context, obj *T) error {
	db := WithContext(ctx, q.DB)
	db = db.Delete(obj)
	if err := db.Error; err != nil {
		return utils.DbError(err)
	}
	return nil
}

func (q *dbQuery[T]) DeleteList(ctx context.Context, obj *[]T) error {
	db := WithContext(ctx, q.DB)
	for _, o := range *obj {
		db = db.Delete(o)
		if err := db.Error; err != nil {
			return utils.DbError(err)
		}
	}
	return nil
}

func (q *dbQuery[T]) DeleteWhere(ctx context.Context, whereCauses *[]dbre.WhereCause) error {
	var obj T
	db := WithContext(ctx, q.DB)

	addWhere(db, whereCauses)

	db = db.Delete(&obj)
	if err := db.Error; err != nil {
		return utils.DbError(err)
	}
	return nil
}

func WithContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	return db.WithContext(ctx)
}

func addWhere(db *gorm.DB, whereCauses *[]dbre.WhereCause) {
	if whereCauses != nil {
		for _, w := range *whereCauses {
			if w.Type == dbre.And {
				db.Where(w.Query, w.Args...)
			} else {
				db.Or(w.Query, w.Args...)
			}
		}
	}
}
