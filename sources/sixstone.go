package sources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WALL-EEEEEEE/gagdets/core"
	"github.com/bobg/go-generics/set"
	log "github.com/sirupsen/logrus"

	"github.com/gocolly/colly/v2"
)

var sixtone_news_api = "https://api.sixthtone.com/cont/nodeCont/getByNodeId"

type SixtoneSpider struct {
	core.Spider
	cnt  int
	page int
}

type NewsPost struct {
	NodeId   string `json:"nodeId"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

func NewNewsPost(page int) NewsPost {
	return NewsPost{NodeId: "26166", PageNum: page, PageSize: 20}
}

func NewSixtoneSpider() SixtoneSpider {
	urls := []string{sixtone_news_api}
	spider := SixtoneSpider{core.NewSpider("SixtoneSpider", urls), 0, 1}
	return spider
}

func (spider *SixtoneSpider) Run(collector *core.Collector) {
	c := colly.NewCollector()
	urls := set.New(spider.Urls...)
	/*
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			req_url := e.Request.URL.String()
			href := e.Attr("href")
			news_href_regex := regexp.MustCompile(`\/news\/.*`)
			if news_href_regex.MatchString(req_url) {
				item_title := e.Text
				item_url := fmt.Sprintf("https://www.sixthtone.com%s", href)
				log.Infof("Find an Item: %s (%s)", item_title, item_url)
				e.Request.Visit(item_url)
			}
		})
	*/
	c.OnResponse(func(r *colly.Response) {
		nodeId := r.Ctx.Get("nodeId")
		if r.StatusCode != 200 {
			log.Warnf("Failed to get %s (Request Exception: %s)!", nodeId, r.Body)
			return
		}
		var postListRsp interface{}
		err := json.Unmarshal(r.Body, postListRsp)
		if err != nil {
			log.Warnf("Failed to get %s (Json Parse Exception: %s)!", nodeId, r.Body)
			return
		}
		log.Infof("%+v", postListRsp)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	for _, url := range urls.Slice() {
		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(NewNewsPost(1))
		if err != nil {
			log.Fatal(err)
		}
		ctx := colly.NewContext()
		headers := http.Header{
			"Accept":       []string{"application/json"},
			"Content-Type": []string{"application/json"},
		}
		ctx.Put("nodeId", "26166")
		c.Request("POST", url, b, ctx, headers)
	}
}

func init() {
	sixtone_spider := NewSixtoneSpider()
	core.Reg.Register(&sixtone_spider)
}
