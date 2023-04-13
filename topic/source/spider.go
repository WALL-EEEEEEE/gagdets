package source

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

var site_url string = "https://www.sixthtone.com/"

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit(site_url)
}
