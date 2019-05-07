package main

import (
	"github.com/shmy/crawler/core/downloader"
	"github.com/shmy/crawler/core/engine"
	"github.com/shmy/crawler/example/zuidazy_xml/pipeline"
	"github.com/shmy/crawler/example/zuidazy_xml/processer"
)

func main() {
	p := &processer.Zuidazy{}
	urls := make([]engine.UrlWithParams, 0)
	for i := 3; i < 5; i++ {
		urls = append(urls, engine.UrlWithParams{
			Url:    p.GetListUrl(i, 1),
			Params: i,
		})
	}
	engine.NewEngine(&processer.Zuidazy{}).
		PutUrls(urls, downloader.TEXT).
		SetPipeline(pipeline.NewFilePipeline()).
		SetThreadnum(100).
		Run()
}
