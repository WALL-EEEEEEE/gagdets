package sources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	. "github.com/WALL-EEEEEEE/gagdets/core"
	"github.com/bobg/go-generics/set"
	log "github.com/sirupsen/logrus"

	"github.com/gocolly/colly/v2"
)

var sixtone_news_api = "https://api.sixthtone.com/cont/nodeCont/getByNodeId"

type SixtoneSpider struct {
	Spider
	cnt  int
	page int
}

type NewsPost struct {
	NodeId   string `json:"nodeId"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

func NewNewsPost(nodeId string, page int) NewsPost {
	return NewsPost{NodeId: string(nodeId), PageNum: page, PageSize: 20}
}

func NewSixtoneSpider() SixtoneSpider {
	urls := []string{
		"26166,News",          //<News> section in SixTone
		"26167,Features",      //<Features> section in SixTone
		"26168,VOICE &OPNION", //<VOICES & OPINION> section in SixTone
		"26290,MULTIMEDIA",    //<MULTIMEDIA> section in SixTone
		"26291,DAILY TONES",   //<DAILY TONES> section in SixTone
		"26301,SIXTH TONE ×",  //<SIXTH TONE ×> section in SixTone
	}
	spider := SixtoneSpider{NewSpider("SixtoneSpider", urls), 0, 1}
	return spider
}

func pagingNews(node_name string, node_id string, page int, c *colly.Collector) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(NewNewsPost(node_id, page))
	if err != nil {
		log.Fatal(err)
	}
	ctx := colly.NewContext()
	headers := http.Header{
		"Accept":       []string{"application/json"},
		"Content-Type": []string{"application/json"},
	}
	log.Infof("Getting <%s>(%s) Section ... at %d page", node_name, node_id, page)
	ctx.Put("nodeId", node_id)
	ctx.Put("nodeName", node_name)
	ctx.Put("page", page)
	c.Request("POST", sixtone_news_api, bytes.NewReader(b.Bytes()), ctx, headers)
}

func (spider *SixtoneSpider) Run(collector *Collector) {
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
	//c.SetProxy("http://127.0.0.1:8889")
	c.OnResponse(func(r *colly.Response) {
		nodeId := r.Ctx.Get("nodeId")
		nodeName := r.Ctx.Get("nodeName")
		page, _ := strconv.Atoi(r.Ctx.Get("page"))

		if r.StatusCode != 200 {
			log.Warnf("Failed to get %s (Request Exception: %s)!", nodeId, r.Body)
			return
		}
		var postListRsp map[string]interface{}
		err := json.Unmarshal(r.Body, &postListRsp)
		if err != nil {
			log.Warnf("Failed to get %s (Json Parse Exception: %s)!", nodeId, r.Body)
			return
		}
		ret_status := postListRsp["code"]
		if int(ret_status.(float64)) != 200 {
			log.Warnf("Failed to get %s (Response invalid: %+v)!", nodeId, postListRsp)
			return
		}
		news, _ := postListRsp["data"].(map[string]interface{})["pageInfo"].(map[string]interface{})["list"].([]interface{})
		has_next, _ := postListRsp["data"].(map[string]interface{})["pageInfo"].(map[string]interface{})["hasNext"].(bool)
		//_total, _ := postListRsp["data"].(map[string]interface{})["pageInfo"].(map[string]interface{})["total"].(float64)
		if len(news) < 1 {
			log.Warnf("%s News is empty!", nodeId)
			return
		}
		for _, new := range news {
			title := new.(map[string]interface{})["name"].(string)
			ptime := new.(map[string]interface{})["pubTimeLong"]
			id := new.(map[string]interface{})["contId"]
			if id == nil {
				continue
			}
			id = int(id.(float64))
			time_f := ""
			if ptime != nil {
				time_f = time.UnixMilli(int64(ptime.(float64))).Format("2006-01-02 15:04:05")
			}
			url := fmt.Sprintf("https://www.sixthtone.com/news/%d", id)
			log.Infof("Get an item:  %s (url: %s, pub: %s)", title, url, time_f)
		}
		if has_next {
			page = page + 1
			pagingNews(nodeName, nodeId, page, c)
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Errorf("Failed to get Request: %s (Exception: %s) ", r.Request.URL.String(), err)
	})
	for _, nodeId := range urls.Slice() {
		node_splits := strings.Split(nodeId, ",")
		node_id := node_splits[0]
		node_name := node_splits[1]
		pagingNews(node_name, node_id, 1, c)
	}
}

func init() {
	sixtone_spider := NewSixtoneSpider()
	Reg.Register(&sixtone_spider)
}
