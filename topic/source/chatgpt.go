package source

import (
	"fmt"
)

type ChatGPT struct{}

func (robot *ChatGPT) Run() {
	fmt.Printf("robot: %v\n", robot)
}
