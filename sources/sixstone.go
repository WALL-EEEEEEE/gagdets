package sources

import (
	"encoding/json"
	"fmt"

	"github.com/WALL-EEEEEEE/gagdets/core"
	"github.com/bobg/go-generics/set"
	log "github.com/sirupsen/logrus"

	"github.com/gocolly/colly/v2"
)

var sixtone_news_api = "https://api.sixthtone.com/cont/nodeCont/getByNodeId"

type SixtoneSpider struct {
	core.Spider
	cnt int
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
	spider := SixtoneSpider{core.NewSpider("SixtoneSpider", urls), 0}
	return spider
}

func (spider *SixtoneSpider) Run(collector *core.Collector) {
	c := colly.NewCollector()
	urls := set.New(spider.Urls...)

	c.OnResponse(func(r *colly.Response) {
		log.Info(r.StatusCode)
		log.Info(r.Body)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	for _, url := range urls.Slice() {
		data, err := json.Marshal(NewNewsPost(1))
		if err != nil {
			log.Fatal(err)
		}
		c.PostRaw(url, data)
	}
}

func init() {
	sixtone_spider := NewSixtoneSpider()
	core.Exec.Add(&sixtone_spider)
}
