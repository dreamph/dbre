package repomodels

import (
	"github.com/dreamph/dbre"
	"github.com/dreamph/dbre/example/core/models"
)

type CountryListCriteria struct {
	Status int32             `json:"status" example:"20"`
	Limit  *models.PageLimit `json:"limit"`
	Sort   *dbre.Sort        `json:"sort"`
}
