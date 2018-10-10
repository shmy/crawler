package page

import "crawler/core/common/request"

func NewPage(req *request.Request) *Page {
	return &Page{
		req: req,
	}
}

type Page struct {
	body     string
	req      *request.Request
	requests []*request.Request
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
