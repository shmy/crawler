package pipeline

import (
	"crawler/core/common/page_items"
	"fmt"
)

func NewPipelineConsole() *PipelineConsole {
	return &PipelineConsole{}
}

var (
	n = 0
)

type PipelineConsole struct{}

func (p *PipelineConsole) Process(items *page_items.PageItems) {
	t := items.GetItems()
	if len(t) > 0 {
		n++
		fmt.Println("Get Data: ", t, n)
	}
}
