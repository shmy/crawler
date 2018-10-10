package request

func NewRequest(url string, respType int) *Request {
	return &Request{
		url,
		respType,
	}
}

type Request struct {
	url      string
	respType int
}

func (r *Request) GetUrl() string {
	return r.url
}

func (r *Request) GetRespType() int {
	return r.respType
}
