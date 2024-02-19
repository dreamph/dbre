package dbre

import (
	errs "github.com/pkg/errors"

	"fmt"
)

const (
	ASC  = "ASC"
	DESC = "DESC"
)

type Sort struct {
	SortBy        string `json:"sortBy"`
	SortDirection string `json:"sortDirection" example:"DESC" enums:"ASC,DESC"`
}

type SortParam struct {
	SortFieldMapping map[string]string
	Sort             *Sort
	DefaultSort      *Sort
}

func SortSQL(param *SortParam) (string, error) {
	if param == nil {
		return "", errs.New("required sortParam")
	}
	if param.SortFieldMapping == nil {
		return "", errs.New("required SortFieldMapping")
	}

	if param.Sort == nil && param.DefaultSort == nil {
		return "", nil
	}

	if param.Sort == nil {
		return fmt.Sprintf("%s %s", param.SortFieldMapping[param.DefaultSort.SortBy], param.DefaultSort.SortDirection), nil
	}

	if param.Sort.SortBy == "" {
		return fmt.Sprintf("%s %s", param.SortFieldMapping[param.DefaultSort.SortBy], param.DefaultSort.SortDirection), nil
	}

	dbField, ok := param.SortFieldMapping[param.Sort.SortBy]
	if !ok {
		return "", errs.New("sortBy not support :" + param.Sort.SortBy)
	}

	return fmt.Sprintf("%s %s", dbField, param.Sort.SortDirection), nil
}
