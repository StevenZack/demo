package storage

func NewTable(name string) *Table {
	return &Table{
		Name: name,
	}
}

func (t *Table) Add(r *Row) error {
	if len(t.Pages) == 0 {
		t.createPage(nil, r)
		return nil
	}
	return t.Pages[0].Add(r)
}

func (t *Table) createPage(parent *Page, r *Row) *Page {
	page := &Page{
		Table:    t,
		Offset:   int64(len(t.Pages)),
		Rows:     []*Row{r},
		Parent:   parent,
		Children: []*Page{nil, nil},
	}
	t.Pages = append(t.Pages, page)
	return page
}

func (t *Table) Print()  {
	
}