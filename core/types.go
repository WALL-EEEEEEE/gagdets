package core

import "time"

type Topic struct {
	Content       string `json:`
	CreateTime    *time.Time
	Source        string
	Keywords      string
	Level         string
	Description   string
	TaskStatus    string
	TaskPublishAt string
	Author        string
	ExtraLink     string
}
