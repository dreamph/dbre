package query

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

type whereBuilder struct {
	whereCauses []WhereCause
}

func NewWhereBuilder() *whereBuilder {
	return &whereBuilder{}
}

func (w *whereBuilder) Where(query string, args ...interface{}) *whereBuilder {
	w.whereCauses = append(w.whereCauses, WhereCause{
		Type:  And,
		Query: query,
		Args:  args,
	})
	return w
}

func (w *whereBuilder) WhereOr(query string, args ...interface{}) *whereBuilder {
	w.whereCauses = append(w.whereCauses, WhereCause{
		Type:  Or,
		Query: query,
		Args:  args,
	})
	return w
}

func (w *whereBuilder) WhereCauses() *[]WhereCause {
	return &w.whereCauses
}
