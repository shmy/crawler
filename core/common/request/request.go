package request

func NewRequest(url string, params interface{}, respType int) *Request {
	return &Request{
		url,
		params,
		respType,
	}
}

type Request struct {
	url      string
	params   interface{}
	respType int
}

func (r *Request) GetUrl() string {
	return r.url
}
func (r *Request) GetParams() interface{} {
	return r.params
}
func (r *Request) GetRespType() int {
	return r.respType
}
