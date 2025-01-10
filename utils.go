package dbre

const (
	Or  = 2
	And = 1
)

type Limit struct {
	Offset   int64 `json:"offset"`
	PageSize int64 `json:"pageSize"`
}

type WhereCause struct {
	Type  int
	Query string
	Args  []interface{}
}

type WhereBuilder struct {
	whereCauses []WhereCause
}

func NewWhereBuilder() *WhereBuilder {
	return &WhereBuilder{}
}

func (w *WhereBuilder) Where(query string, args ...interface{}) *WhereBuilder {
	w.whereCauses = append(w.whereCauses, WhereCause{
		Type:  And,
		Query: query,
		Args:  args,
	})
	return w
}

func (w *WhereBuilder) WhereOr(query string, args ...interface{}) *WhereBuilder {
	w.whereCauses = append(w.whereCauses, WhereCause{
		Type:  Or,
		Query: query,
		Args:  args,
	})
	return w
}

func (w *WhereBuilder) WhereCauses() *[]WhereCause {
	return &w.whereCauses
}
