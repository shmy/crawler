package request

func NewRequest(url string, params interface{}, respType int, maxRetryCount int) *Request {
	return &Request{
		url,
		params,
		respType,
		0,
		maxRetryCount,
	}
}

type Request struct {
	url               string
	params            interface{}
	respType          int
	currentRetryCount int
	maxRetryCount     int
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
func (r *Request) GetCurrentRetryCount() int {
	return r.currentRetryCount
}
func (r *Request) GetMaxRetryCount() int {
	return r.maxRetryCount
}
func (r *Request) AddCurrentRetryCount() {
	r.currentRetryCount += 1
}
