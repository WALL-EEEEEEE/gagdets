package sources

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	. "github.com/WALL-EEEEEEE/gagdets/core"
	. "github.com/WALL-EEEEEEE/gagdets/items"
	"github.com/WALL-EEEEEEE/gagdets/utils"
	"github.com/bobg/go-generics/set"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

var iets_site = "https://ieltsonlinetests.com/speaking-test-collection"

type IETSOnlineTestsSpider struct {
	Spider
	cnt int
}

func NewIETSSpider() IETSOnlineTestsSpider {
	urls := []string{iets_site}
	spider := IETSOnlineTestsSpider{NewSpider("IETSOnlineTestsSpider", urls), 0}
	return spider
}

func (spider *IETSOnlineTestsSpider) parseTopicList(e *colly.HTMLElement) {
	part1 := e.ChildTexts("#recording-accordion > div:nth-child(2) > #part1 > div > ul > li") //Introduction and Interview Part
	part2 := e.ChildTexts("#recording-accordion > div:nth-child(3) > #part2 > div > ul > li") //Topic Part
	part3 := e.ChildTexts("#recording-accordion > div:nth-child(4) > #part3 > div > ul > li") //Topic Discussion
	stripHtml := func(s string) string {
		r := regexp.MustCompile(`<.*?>`)
		return r.ReplaceAllString(s, "")
	}

	for _, item := range utils.Concat([][]string{part1, part2, part3}) {
		item = strings.TrimSpace(stripHtml(item))
		topic := Topic{
			Content:    item,
			Source:     "IELTS",
			CreateTime: time.Now(),
			TaskStatus: "未使用，可充当今日任务",
			Author:     "Wall'E's Robot",
			ExtraLink:  e.Request.URL.String(),
		}
		log.Infof("Topic: %s", item)
		spider.GetOutputStream().In() <- topic
		spider.cnt++
	}
}

func (spider *IETSOnlineTestsSpider) Run() {
	c := colly.NewCollector()
	urls := set.New(spider.Urls...)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		req_url := e.Request.URL.String()
		href := e.Attr("href")
		topic_href_regex := regexp.MustCompile("/.*speaking-practice-test.*")
		page_url_regex := regexp.MustCompile(`https:\/\/ieltsonlinetests\.com\/speaking\-test\-collection(\?page=\d+)?`)
		log.Debugf("%s, %s: %v, %v", req_url, href, (urls.Has(req_url) || page_url_regex.MatchString(req_url)), topic_href_regex.MatchString(href))
		if (urls.Has(req_url) || page_url_regex.MatchString(req_url)) && topic_href_regex.MatchString(href) {
			topic_name := e.Text
			log.Debugf("Find a topic: %s", topic_name)
			topic_url := fmt.Sprintf("https://ieltsonlinetests.com%s", href)
			log.Debug(topic_url)
			e.Request.Visit(topic_url)
		}
	})
	c.OnHTML("div[id='recording-accordion']", func(h *colly.HTMLElement) {
		spider.parseTopicList(h)
	})
	c.OnHTML("nav[class~='pager']", func(h *colly.HTMLElement) {
		next_page_href := h.ChildAttr("ul > li.pager__item--next > a", "href")
		if len(next_page_href) > 0 {
			next_page_link := "https://ieltsonlinetests.com/speaking-test-collection" + next_page_href
			log.Debug(next_page_link)
			h.Request.Visit(next_page_link)
		} else {
			log.Info("There are'nt any new pages anymore!")
		}
	})
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Infof("Failed to request information of IETS topics (Request Error: %s)!", r.Body)
		}
	})
	for _, url := range spider.Urls {
		c.Visit(url)
	}
	c.Wait()
	log.Infof("Spider found: %d topics.", spider.cnt)
}

func init() {
	iets_spider := NewIETSSpider()
	Reg.Register(&iets_spider)
}
