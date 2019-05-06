package scheduler

import "github.com/shmy/crawler/core/common/request"

type Scheduler interface {
	Put(req *request.Request)
	Get() *request.Request
	Count() int
}
