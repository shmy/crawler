package page

import (
	"github.com/shmy/crawler/core/common/page_items"
	"github.com/shmy/crawler/core/common/request"
)

func NewPage(req *request.Request) *Page {
	return &Page{
		req:       req,
		pageItems: page_items.NewPageItems(),
	}
}

type Page struct {
	body      string
	pageItems *page_items.PageItems
	req       *request.Request
	requests  []*request.Request
}

// 设置Body
func (p *Page) SetBodyString(body string) *Page {
	p.body = body
	return p
}

// 获取Body
func (p *Page) GetBodyString() string {
	return p.body
}

// 添加一个Request
func (p *Page) PutRequest(req *request.Request) *Page {
	p.requests = append(p.requests, req)
	return p
}

// 添加多个个Request
func (p *Page) PutRequests(reqs []*request.Request) *Page {
	for _, req := range reqs {
		p.PutRequest(req)
	}
	return p
}

// 获取Requests
func (p *Page) GetRequests() []*request.Request {
	return p.requests
}

// 获取本尊的req
func (p *Page) GetRequest() *request.Request {
	return p.req
}

// 添加结果
func (p *Page) PutField(key string, value interface{}) *Page {
	p.pageItems.PutItem(key, value)
	return p
}

func (p *Page) GetPageItems() *page_items.PageItems {
	return p.pageItems
}
