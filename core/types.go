package core

import "time"

type Topic struct {
	Content    string
	CreateTime *time.Time
	Source     string
}
