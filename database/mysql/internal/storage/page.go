package storage

import (
	"errors"
	sync "sync"

	"github.com/StevenZack/tools/alg"
)

const (
	PAGE_SIZE     = 16 << 10
	PAGE_MAX_ROWS = 4
)

func (p *Page) Len() int {
	return len(p.Rows)
}

var pageLocker sync.Mutex

func (p *Page) Add(r *Row) error {
	for i, row := range p.Rows {
		if row.Key() == r.Key() {
			return errors.New("duplicated key :" + r.Key())
		}
		if r.Key() < row.Key() {
			if p.Children[i] == nil {
				p.insertRow(i, r)
				return nil
			}
			// todo
			return p.Children[i].Add(r)
		}
	}
	if p.Children[len(p.Rows)] == nil {
		p.insertRow(len(p.Rows), r)
		return nil
	}

	return p.Children[len(p.Rows)].Add(r)
}

func (p *Page) insertRow(i int, r *Row) {
	pageLocker.Lock()
	p.Rows = append(p.Rows[:i], append([]*Row{r}, p.Rows[i:]...)...)
	p.Children = append(p.Children[:i], append([]*Page{nil}, p.Children[1:]...)...)
	pageLocker.Unlock()
}

func (p *Page) Child(i int) alg.ITreeNode {
	return p.Children[i]
}

func (p *Page) ChildNum() int {
	return len(p.Children)
}

// func (p *Page) String() string {
// 	ss := []string{}
// 	for _, r := range p.Rows {
// 		ss = append(ss, r.Key())
// 	}
// 	return "{" + strings.Join(ss, ",") + "}"
// }
