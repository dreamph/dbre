package repository

import (
	"context"

	"github.com/dreamph/dbre/query"
)

type CountryRepository interface {
	WithTx(db *query.AppIDB) CountryRepository

	Create(ctx context.Context, obj *domain.Country) (*domain.Country, error)
	Update(ctx context.Context, obj *domain.Country) (*domain.Country, error)
	UpdateForce(ctx context.Context, obj *domain.Country) (*domain.Country, error)
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*domain.Country, error)
	FindOne(ctx context.Context, obj *domain.Country) (*domain.Country, error)

	List(ctx context.Context, obj *repomodels.CountryListCriteria) (*[]domain.Country, int64, error)
}

type countryRepository struct {
	query query.DB[domain.Country]
}

func NewCountryRepository(db *query.AppIDB) CountryRepository {
	return &countryRepository{
		query: bun.New[domain.Country](db),
	}
}

func (r *countryRepository) WithTx(tx *query.AppIDB) CountryRepository {
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
	whereBuilder := query.NewWhereBuilder()

	if obj.Status != 0 {
		whereBuilder.Where("status = ?", obj.Status)
	}

	whereCauses := whereBuilder.WhereCauses()
	total, err := r.query.CountWhere(ctx, whereCauses)
	if err != nil {
		return nil, 0, err
	}
	if total > 0 {
		sortSQL, err := database.SortSQL(&database.SortParam{
			SortFieldMapping: map[string]string{
				"id":     "id",
				"nameEn": "name_en",
				"nameTh": "name_th",
				"code":   "code",
				"status": "status",
				//"state":             "state",
				//"createBy":   "create_by",
				//"createDate": "create_date",
				//"updateBy":   "update_by",
				//"updateDate": "update_date",
			},
			Sort: obj.Sort,
			DefaultSort: &models.Sort{
				SortBy:        "nameEn",
				SortDirection: "DESC",
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
