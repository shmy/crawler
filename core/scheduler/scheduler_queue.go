package scheduler

import (
	"container/list"
	"github.com/shmy/crawler/core/common/request"
	"sync"
)

func NewQueueScheduler() *QueueScheduler {
	queue := list.New()
	locker := &sync.Mutex{}
	return &QueueScheduler{
		queue:  queue,
		locker: locker,
	}
}

type QueueScheduler struct {
	locker *sync.Mutex
	queue  *list.List
}

func (q *QueueScheduler) Put(req *request.Request) {
	defer q.locker.Unlock()
	q.locker.Lock()
	q.queue.PushBack(req)
}

func (q *QueueScheduler) Get() *request.Request {
	defer q.locker.Unlock()
	q.locker.Lock()
	if q.queue.Len() <= 0 {
		return nil
	}
	e := q.queue.Front()
	req := e.Value.(*request.Request)
	q.queue.Remove(e)
	return req
}

func (q *QueueScheduler) Count() int {
	defer q.locker.Unlock()
	q.locker.Lock()
	length := q.queue.Len()
	return length
}
