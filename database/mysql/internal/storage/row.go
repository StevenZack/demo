package storage

func NewRow(values []string) *Row {
	return &Row{
		Values: values,
	}
}

func (r *Row) Key() string {
	return r.Values[0]
}
