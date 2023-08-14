package items

type NewsItem struct {
	Content string `json:"新闻内容"`
	PubTime string `json:"发布时间"`
	Url     string `json:"新闻URL"`
	Title   string `json:"新闻标题"`
}
