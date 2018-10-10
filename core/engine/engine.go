package engine

import (
	"crawler/core/common/request"
	"crawler/core/common/resource_manage"
	"crawler/core/downloader"
	"crawler/core/processer"
	"crawler/core/scheduler"
	"time"
)

func NewEngine(p processer.Processer) *Engine {
	return &Engine{
		processer:  p,
		downloader: downloader.NewDownloaderHttp(),
		scheduler:  scheduler.NewQueueScheduler(true),
		threadnum:  10,
	}
}

type Engine struct {
	processer  processer.Processer
	downloader downloader.Downloader
	scheduler  scheduler.Scheduler
	rm         resource_manage.ResourceManage
	threadnum  int
}

// 添加一个url
func (e *Engine) AddUrl(url string, respType int) *Engine {
	e.scheduler.Put(request.NewRequest(url, respType))
	return e
}

// 添加多个url
func (e *Engine) AddUrls(urls []string, respType int) *Engine {
	for _, url := range urls {
		e.AddUrl(url, respType)
	}
	return e
}

// 设置线程er
func (e *Engine) SetThreadnum(num int) *Engine {
	e.threadnum = num
	return e
}

// 运行
func (e *Engine) Run() {
	e.rm = resource_manage.NewResourceManageChan(e.threadnum)
	for {
		req := e.scheduler.Get()
		// 没有数据 并且没有req
		if !e.rm.Has() && req == nil {
			break
		}
		if req == nil {
			// 队列没有数据 休息500ms
			time.Sleep(300 * time.Millisecond)
			continue
		}
		// 一直往缓冲chan里送数据 送满了 没人收 就会等待 从而卡住for循环
		e.rm.GetOne()
		go func() {
			// 消费掉一个chan
			defer e.rm.FreeOne()
			p := e.downloader.Download(req)
			// 下载出错
			if p == nil {
				// 重新加入队列
				e.scheduler.Put(req)
				return
			}
			// 先进行处理
			e.processer.Process(p)
			// 然后获取新的requests
			r := p.GetRequests()
			// 加入队列
			e.putRequests(r)
		}()
	}
}

// 添加request
func (e *Engine) putRequests(reqs []*request.Request) {
	if len(reqs) < 1 {
		return
	}
	for _, req := range reqs {
		e.scheduler.Put(req)
	}
}
