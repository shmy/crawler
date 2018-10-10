package scheduler

import "crawler/core/common/request"

type Scheduler interface {
	Put(req *request.Request)
	Get() *request.Request
	Count() int
}
