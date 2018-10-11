package page_items

type Item map[string]interface{}

func NewPageItems() *PageItems {
	return &PageItems{
		items: make(Item),
	}
}

type PageItems struct {
	items Item
}

func (p *PageItems) PutItem(key string, value interface{}) {
	p.items[key] = value
}
func (p *PageItems) GetItem(key string) interface{} {
	t, err := p.items[key]
	if err != true {
		return nil
	}
	return t
}
func (p *PageItems) GetAllItem() Item {
	return p.items
}
