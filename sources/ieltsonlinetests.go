package sources

import (
	"fmt"
	"regexp"
	"time"

	"github.com/WALL-EEEEEEE/gagdets/core"
	"github.com/WALL-EEEEEEE/gagdets/utils"
	"github.com/bobg/go-generics/set"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

var iets_site = "https://ieltsonlinetests.com/speaking-test-collection"

type IETSOnlineTestsSpider struct {
	core.Spider
}

func NewIETSSpider() IETSOnlineTestsSpider {
	urls := []string{iets_site}
	spider := IETSOnlineTestsSpider{core.NewSpider("IETSOnlineTestsSpider", urls)}
	return spider
}

func (spider *IETSOnlineTestsSpider) parseTopicList(e *colly.HTMLElement) (topics []*core.Topic) {
	part1 := e.ChildTexts("#recording-accordion > div:nth-child(2) > #part1 > div > ul > li") //Introduction and Interview Part
	part2 := e.ChildTexts("#recording-accordion > div:nth-child(3) > #part2 > div > ul > li") //Topic Part
	part3 := e.ChildTexts("#recording-accordion > div:nth-child(4) > #part3 > div > ul > li") //Topic Discussion

	for _, item := range utils.Concat([][]string{part1, part2, part3}) {
		topic := core.Topic{
			Content:    item,
			Source:     "IELTS",
			CreateTime: time.Now(),
			TaskStatus: "未使用，可充当今日任务",
			Author:     "Wall'E's Robot",
			ExtraLink:  e.Request.URL.String(),
		}
		topics = append(topics, &topic)
	}
	return topics
}

func (spider *IETSOnlineTestsSpider) Run() {
	c := colly.NewCollector()
	urls := set.New(spider.Urls...)

	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		req_url := e.Request.URL.String()
		topic_href := e.Attr("href")
		topic_href_regex := regexp.MustCompile("/ielts-mock-test.+-speaking-practice-test.*")
		if urls.Has(req_url) && topic_href_regex.MatchString(topic_href) {
			topic_name := e.Text
			log.Infof("Find a topic: %s", topic_name)
			topic_url := fmt.Sprintf("https://ieltsonlinetests.com%s", topic_href)
			log.Info(topic_url)
			e.Request.Visit(topic_url)
		}
	})
	c.OnHTML("div[id='recording-accordion']", func(h *colly.HTMLElement) {
		spider.parseTopicList(h)
	})
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Infof("Failed to request information of IETS topics (Request Error: %s)!", r.Body)
		}
	})
	for _, url := range spider.Urls {
		c.Visit(url)
	}
}

func init() {
	iets_spider := NewIETSSpider()
	core.Exec.Add(&iets_spider)
}
