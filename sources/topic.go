package sources

import "time"

type Topic struct {
	Content       string    `json:"话题内容"`
	CreateTime    time.Time `json:"创建时间"`
	Source        string    `json:"话题出处"`
	Keywords      string    `json:"话题关键字"`
	Level         string    `json:"话题难度"`
	Description   string    `json:"话题辅助描述"`
	TaskStatus    string    `json:"任务使用状态"`
	TaskPublishAt string    `json:"任务布置时间"`
	Author        string    `json:"出题人"`
	ExtraLink     string    `json:"话题延生外链"`
}
