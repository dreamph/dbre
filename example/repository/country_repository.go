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
