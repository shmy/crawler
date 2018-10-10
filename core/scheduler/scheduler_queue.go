package scheduler

import (
	"container/list"
	"crawler/core/common/request"
	"crypto/md5"
	"sync"
)

func NewQueueScheduler(rmDuplicate bool) *QueueScheduler {
	queue := list.New()
	rmKey := make(map[[md5.Size]byte]*list.Element)
	locker := &sync.Mutex{}
	return &QueueScheduler{rm: rmDuplicate, queue: queue, rmKey: rmKey, locker: locker}
}

type QueueScheduler struct {
	locker *sync.Mutex
	rm     bool
	rmKey  map[[md5.Size]byte]*list.Element
	queue  *list.List
}

func (q *QueueScheduler) Put(req *request.Request) {
	q.locker.Lock()
	var key [md5.Size]byte
	if q.rm {
		key = md5.Sum([]byte(req.GetUrl()))
		if _, ok := q.rmKey[key]; ok {
			q.locker.Unlock()
			return
		}
	}
	e := q.queue.PushBack(req)
	if q.rm {
		q.rmKey[key] = e
	}
	q.locker.Unlock()
}

func (q *QueueScheduler) Get() *request.Request {
	q.locker.Lock()
	if q.queue.Len() <= 0 {
		q.locker.Unlock()
		return nil
	}
	e := q.queue.Front()
	req := e.Value.(*request.Request)
	key := md5.Sum([]byte(req.GetUrl()))
	q.queue.Remove(e)
	if q.rm {
		delete(q.rmKey, key)
	}
	q.locker.Unlock()
	return req
}

func (q *QueueScheduler) Count() int {
	q.locker.Lock()
	length := q.queue.Len()
	q.locker.Unlock()
	return length
}
