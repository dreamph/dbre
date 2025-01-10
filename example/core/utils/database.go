package utils

import (
	"github.com/dreamph/dbre"
	"github.com/dreamph/dbre/example/core/models"
)

func ToQueryLimit(pageLimit *models.PageLimit) *dbre.Limit {
	if pageLimit == nil {
		return nil
	}
	limit := &dbre.Limit{}
	limit.PageSize = pageLimit.PageSize
	limit.Offset = (pageLimit.PageNumber - 1) * pageLimit.PageSize
	return limit
}
