package storage

func (r *Row) Key() string {
	return r.Values[0]
}
