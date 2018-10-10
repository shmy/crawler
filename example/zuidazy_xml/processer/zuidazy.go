package processer

import (
	"crawler/core/common/page"
	"crawler/core/common/request"
	"crawler/example/zuidazy_xml/entity"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Zuidazy struct{}

var (
	i = 0
)

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
		respType = p.GetRequest().GetRespType()
	)
	r := &entity.Rss{}
	err := xml.Unmarshal([]byte(p.GetBodyString()), &r)
	if err != nil {
		p.PutRequest(request.NewRequest(url, respType))
		return
	}
	if len(r.List.Videos) < 1 {
		return
	}
	for _, item := range r.List.Videos {
		p.PutRequest(request.NewRequest(z.getDetailUrl(item.Id), respType))
	}
	p.PutRequest(request.NewRequest(z.getNextPageUrl(url), respType))
}

func (z *Zuidazy) parserDetail(p *page.Page) {
	var (
		url      = p.GetRequest().GetUrl()
		respType = p.GetRequest().GetRespType()
	)
	r := &entity.Rss{}
	err := xml.Unmarshal([]byte(p.GetBodyString()), &r)
	if err != nil {
		p.PutRequest(request.NewRequest(url, respType))
		return
	}
	if len(r.List.Videos) < 1 {
		return
	}
	i++
	video := r.List.Videos[0]
	fmt.Println(video.Name, i)
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
