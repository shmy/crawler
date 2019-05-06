package pipeline

import (
	"fmt"
	"github.com/shmy/crawler/core/common/page_items"
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
