package processer

import "github.com/shmy/crawler/core/common/page"

type Processer interface {
	Process(p *page.Page)
}
