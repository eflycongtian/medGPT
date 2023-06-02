package model

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"sinohealth.com/medGPT/conf"
)

var client *openai.Client

func init() {
	apiKey := conf.ApplicationConfig.String("openai.apikey")
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.openai-proxy.com/v1"
	client = openai.NewClientWithConfig(config)
}

func CreateChatCompletion(content string) (resp openai.ChatCompletionResponse, err error) {
	resp, err = client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo0301,
			Temperature: 0.5,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	return resp, err
}
