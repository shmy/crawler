package downloader

import (
	"github.com/shmy/crawler/core/common/page"
	"github.com/shmy/crawler/core/common/request"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func NewDownloaderHttp() *DownloaderHttp {
	return &DownloaderHttp{}
}

type DownloaderHttp struct{}

func (d *DownloaderHttp) Download(req *request.Request) *page.Page {
	p := page.NewPage(req)
	switch req.GetRespType() {
	case TEXT:
		return d.downloadText(req, p)
	default:
		log.Println("error request type: ", req.GetRespType())
	}
	return p
}

// 下载纯文本
func (d *DownloaderHttp) downloadText(request *request.Request, p *page.Page) *page.Page {
	client := &http.Client{
		// 超时10s
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", request.GetUrl(), nil)
	if err != nil {
		//log.Println("NewRequest Error: ", err.Error())
		return nil
	}
	//req.Header.Set("User-Agent", randomUA())
	res, err := client.Do(req)
	if err != nil {
		//log.Println("Do Request Error: ", err.Error())
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//log.Println("Read Content Error: ", err.Error())
		return nil
	}
	p.SetBodyString(string(body))
	return p
}
