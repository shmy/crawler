package processer

import (
	"encoding/xml"
	"fmt"
	"github.com/shmy/crawler/core/common/page"
	"github.com/shmy/crawler/core/common/request"
	"github.com/shmy/crawler/example/zuidazy_xml/entity"
	"regexp"
	"strconv"
	"strings"
)

type Zuidazy struct{}

func (z *Zuidazy) Process(p *page.Page) {
	if strings.Contains(p.GetRequest().GetUrl(), "?ac=list") {
		z.parserList(p)
	} else {
		z.parserDetail(p)
	}
}
func (z *Zuidazy) parserList(p *page.Page) {
	var (
		url      = p.GetRequest().GetUrl()
		params   = p.GetRequest().GetParams()
		respType = p.GetRequest().GetRespType()
	)
	r := &entity.Rss{}
	err := xml.Unmarshal([]byte(p.GetBodyString()), &r)
	if err != nil {
		p.PutRequest(request.NewRequest(url, params, respType, 3))
		return
	}
	if len(r.List.Videos) < 1 {
		return
	}
	for _, item := range r.List.Videos {
		p.PutRequest(request.NewRequest(z.getDetailUrl(item.Id), params, respType, 3))
	}
	p.PutRequest(request.NewRequest(z.getNextPageUrl(url), params, respType, 3))
}

func (z *Zuidazy) parserDetail(p *page.Page) {
	var (
		url      = p.GetRequest().GetUrl()
		params   = p.GetRequest().GetParams()
		respType = p.GetRequest().GetRespType()
	)
	r := &entity.Rss{}
	err := xml.Unmarshal([]byte(p.GetBodyString()), &r)
	if err != nil {
		p.PutRequest(request.NewRequest(url, params, respType, 3))
		return
	}
	if len(r.List.Videos) < 1 {
		return
	}
	for _, video := range r.List.Videos {
		p.PutField("name", video.Name)
		p.PutField("params", params)
	}
}

func (z *Zuidazy) getNextPageUrl(url string) string {
	reg := regexp.MustCompile(`&pg=(\d+)`)
	s := reg.ReplaceAllStringFunc(url, func(s string) string {
		d := regexp.MustCompile(`(\d+)`).FindString(s)
		n, _ := strconv.Atoi(d)
		n++
		return fmt.Sprintf("&pg=%d", n)
	})
	return s
}
func (z *Zuidazy) GetListUrl(tid, pg int) string {
	return fmt.Sprintf(z.getUrlPrefix()+"?ac=list&t=%d&pg=%d", tid, pg)
}
func (z *Zuidazy) getDetailUrl(id int) string {
	return fmt.Sprintf(z.getUrlPrefix()+"?ac=videolist&ids=%d", id)
}
func (z *Zuidazy) getUrlPrefix() string {
	return "http://www.zdziyuan.com/inc/s_api_m3u8.php"
}
