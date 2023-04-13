package source

import (
	"fmt"

	. "github.com/WALL-EEEEEEE/gagdets/core"

	"github.com/gocolly/colly/v2"
)

type SixtoneSpider struct {
	Spider
}

func (spider *SixtoneSpider) New() {
}

func (spider *SixtoneSpider) run() {
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
