package downloader

import (
	"github.com/shmy/crawler/core/common/page"
	"github.com/shmy/crawler/core/common/request"
)

const (
	TEXT = iota // value --> 0
	HTML        // value --> 1
	JSON        // value --> 2
)

type Downloader interface {
	Download(req *request.Request) *page.Page
}
