package pipeline

import (
	"github.com/shmy/crawler/core/common/page_items"
)

type Pipeline interface {
	Process(*page_items.PageItems)
}
