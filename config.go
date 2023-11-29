package gaddets

type Config struct {
	ChatGPT struct {
		Token string `yml:"token" envconfig:"CHATGPT_TOKEN"`
	} `yml:"chatgpt"`
}
