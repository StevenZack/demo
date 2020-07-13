package storage

import (
	"errors"
	sync "sync"
)

const (
	PAGE_SIZE     = 16 << 10
	PAGE_MAX_ROWS = 4
)

func (p *Page) Len() int {
	return len(p.Rows)
}

var pageLocker sync.Mutex

func (p *Page) Insert(r *Row) error {
	for i, row := range p.Rows {
		if row.Key() == r.Key() {
			return errors.New("duplicated key :" + r.Key())
		}
		if r.Key() < row.Key() {
			if p.Children[i] == nil {
				pageLocker.Lock()
				p.Rows = append(p.Rows[:i], append([]*Row{r}, p.Rows[i:]...)...)
				p.Children = append(p.Children[:i], append([]*Page{nil}, p.Children[1:]...)...)
				pageLocker.Unlock()
				return nil
			}
			// todo
		}
	}
	return nil
}
