package sources

import (
	"fmt"

	"github.com/WALL-EEEEEEE/gagdets/core"
	"github.com/gocolly/colly/v2"
)

type IETSOnlineTestsSpider struct {
	core.Spider
	name string
}

func NewIETSSpider() IETSOnlineTestsSpider {
	spider := IETSOnlineTestsSpider{name: "IETSOnlineTestsSpider"}
	return spider
}

func (spider *IETSOnlineTestsSpider) Run() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	for _, url := range spider.Urls {
		c.Visit(url)
	}
}

func init() {
	iets_spider := NewIETSSpider()
	core.Exec.Add(&iets_spider)
}
