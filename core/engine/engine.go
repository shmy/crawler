package engine

import (
	"crawler/core/common/page"
	"crawler/core/common/request"
	"crawler/core/common/resource_manage"
	"crawler/core/downloader"
	"crawler/core/pipeline"
	"crawler/core/processer"
	"crawler/core/scheduler"
	"log"
	"time"
)

func NewEngine(p processer.Processer) *Engine {
	return &Engine{
		processer:  p,
		downloader: downloader.NewDownloaderHttp(),
		scheduler:  scheduler.NewQueueScheduler(),
		pipeline:   pipeline.NewPipelineConsole(),
		threadnum:  10,
	}
}

type Engine struct {
	processer  processer.Processer
	downloader downloader.Downloader
	scheduler  scheduler.Scheduler
	pipeline   pipeline.Pipeline
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
			time.Sleep(500 * time.Millisecond)
			continue
		}
		// 一直往缓冲chan里送数据 送满了 没人收 就会等待 从而卡住for循环
		e.rm.GetOne()
		go func() {
			// 消费掉一个chan
			defer e.rm.FreeOne()
			e.pageProcess(req)
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

func (e *Engine) pageProcess(req *request.Request) {
	defer func() {
		if err := recover(); err != nil { // do not affect other
			if strErr, ok := err.(string); ok {
				log.Println(strErr)
			} else {
				log.Println("pageProcess error")
			}
		}
	}()
	var p *page.Page
	// 下载页面 如果下载失败 重试三次
	for i := 0; i < 3; i++ {
		p = e.downloader.Download(req)
		if p != nil {
			break
		}
		// 等待5秒重试
		time.Sleep(300 * time.Millisecond)
	}
	// 仍然没有下载完毕
	if p == nil {
		// 重新加入队列
		e.scheduler.Put(req)
		return
	}
	// 先进行处理
	e.processer.Process(p)
	// 然后获取新的requests加入队列
	e.putRequests(p.GetRequests())
	// 处理好的结果交给结果处理函数
	e.pipeline.Process(p.GetPageItems())
}
