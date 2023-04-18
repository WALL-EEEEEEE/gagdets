package core

type Spider struct {
	Runnable
	Urls []string
}

func (*Spider) Run() {
}

func (spider *Spider) New(name string) {
	spider.name = name
}
