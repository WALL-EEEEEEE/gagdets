package sources

import (
	"context"

	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

type ChatGPT struct {
	token    string
	question string
}

func (robot *ChatGPT) Run() {
	client := openai.NewClient(robot.token)
	resp, _ := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: robot.question,
				},
			},
		},
	)
	log.Infof("ChatGPT <-  %v\n", robot.question)
	log.Infof("ChatGPT ->  %v\n", resp)
}
