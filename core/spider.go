package core

type Spider struct {
	Urls []string
}

func (*Spider) Run() {
}

func (spider *Spider) New(urls []string) {
	spider.Urls = urls
}
