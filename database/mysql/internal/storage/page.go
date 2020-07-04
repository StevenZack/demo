package storage

type Page struct {
	data []byte
}

const PAGE_SIZE = 16 << 10 //页固定大小16k

func NewPage() *Page {
	return &Page{
		data: make([]byte, PAGE_SIZE),
	}
}

func (p *Page) FileHeader() []byte {
	return p.data[:38]
}

func (p *Page) PageHeader() []byte {
	return p.data[:56]
}

func (p *Page) Infimum()  {
	
}