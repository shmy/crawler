package engine

import (
	"github.com/shmy/crawler/core/common/request"
	"github.com/shmy/crawler/core/common/resource_manage"
	"github.com/shmy/crawler/core/downloader"
	"github.com/shmy/crawler/core/pipeline"
	"github.com/shmy/crawler/core/processer"
	"github.com/shmy/crawler/core/scheduler"
	"log"
	"time"
)

type UrlWithParams struct {
	Url    string
	Params interface{}
}

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
func (e *Engine) PutUrl(url string, params interface{}, respType int) *Engine {
	e.scheduler.Put(request.NewRequest(url, params, respType))
	return e
}

// 添加多个url
func (e *Engine) PutUrls(urlsWithPrams []UrlWithParams, respType int) *Engine {
	for _, item := range urlsWithPrams {
		e.PutUrl(item.Url, item.Params, respType)
	}
	return e
}

// 设置pipeline
func (e *Engine) SetPipeline(pipeline pipeline.Pipeline) *Engine {
	e.pipeline = pipeline
	return e
}

// 设置downloader
func (e *Engine) SetDownloader(downloader downloader.Downloader) *Engine {
	e.downloader = downloader
	return e
}

// 设置scheduler
func (e *Engine) SetScheduler(scheduler scheduler.Scheduler) *Engine {
	e.scheduler = scheduler
	return e
}

// 设置线程数
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
				log.Println(err)
			}
		}
	}()
	log.Println("Do Get: ", req.GetUrl())
	p := e.downloader.Download(req)
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
