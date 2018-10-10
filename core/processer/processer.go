package processer

import "crawler/core/common/page"

type Processer interface {
	Process(p *page.Page)
}
