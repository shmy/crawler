package main

import (
	"crawler/core/downloader"
	"crawler/core/engine"
	"crawler/processer"
)

func main() {
	p := &processer.Zuidazy{}
	urls := make([]string, 0)
	for i := 3; i < 19; i ++ {
		urls = append(urls, p.GetListUrl(i, 1))
	}
	engine.NewEngine(&processer.Zuidazy{}).
		AddUrls(urls, downloader.TEXT).
		SetThreadnum(10000).
		Run()
}
