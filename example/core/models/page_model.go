package models

// PageLimit ...
type PageLimit struct {
	PageNumber int64 `json:"pageNumber" example:"1"`
	PageSize   int64 `json:"pageSize" example:"10"`
}

// PageResult ...
type PageResult struct {
	PageSize   int64 `json:"pageSize"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	PageNumber int64 `json:"pageNumber"`
}

type PageQuery struct {
	Offset   int64 `json:"offset"`
	PageSize int64 `json:"pageSize"`
}
