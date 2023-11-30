package sources

import (
	"context"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	. "github.com/WALL-EEEEEEE/gagdets"
	. "github.com/WALL-EEEEEEE/gagdets/items"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

type ChatGPT struct {
	Task
	conf *Config
}

func NewChatGPT(conf *Config) ChatGPT {
	task := ChatGPT{Task: NewTask("ChatGPT"), conf: conf}
	task.ITask = &task
	return task
}

func (c *ChatGPT) Run() {
	token := c.conf.ChatGPT.Token
	if len(token) < 1 {
		log.Fatal("ChatGPT need a token supplied!")
	}
	log.Infof("ChatGPT Token: %s", token)
	client := openai.NewClient(c.conf.ChatGPT.Token)
	output_stream := c.GetOutputStream()
	for {
		item, ok := output_stream.Read()
		if !ok {
			break
		}
		question := item.(Topic).Content
		log.Infof("Feed question to ChatGPT:   %s\n", question[:100])
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: question,
					},
				},
			},
		)
		if err != nil {
			log.Error(err)
		}
		log.Infof("Generate topic by ChatGPT:  %+v\n", resp)

	}
}
