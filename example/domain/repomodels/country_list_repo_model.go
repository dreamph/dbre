package repomodels

import (
	"example/core/models"

	"github.com/dreamph/dbre/query"
)

type CountryListCriteria struct {
	Status int32             `json:"status" example:"20"`
	Limit  *models.PageLimit `json:"limit"`
	Sort   *query.Sort       `json:"sort"`
}
