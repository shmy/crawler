package pipeline

import (
	"github.com/shmy/crawler/core/common/page_items"
	"os"
	"path"
)

func NewFilePipeline() *FilePipeline {
	file, err := os.Create(path.Join("1.txt"))
	if err != nil {
		panic(err.Error())
	}
	return &FilePipeline{
		file: file,
	}
}

type FilePipeline struct {
	file *os.File
}

func (f *FilePipeline) Process(items *page_items.PageItems) {
	t := items.GetItems()
	if len(t) > 0 {
		f.file.WriteString(items.GetItem("name").(string) + "\n")
	}
}
