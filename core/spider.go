package core

type Spider struct {
	name string
	Urls []string
}

func (*Spider) Run() {
}

func NewSpider(name string, urls []string) Spider {
	spider := Spider{name: name, Urls: urls}
	return spider
}

func (spider *Spider) GetName() string {
	return spider.name
}
