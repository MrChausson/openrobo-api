package aiclient

import (
	"context"
	"fmt"
	"openrobo-api/src/config"

	"openrobo-api/src/openai"
)

var client *openai.Client
var req openai.ChatCompletionRequest

func init() {
	client = openai.NewClient(config.OpenAIKey)
	req = openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a cute robot companion and show a lot of emotions. End every message with a smiley face representing the emotion you are feeling.",
			},
		},
	}
}

func Ask(prompt string) (string, error) {
	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "ChatCompletion error", err
	}
	req.Messages = append(req.Messages, resp.Choices[0].Message)
	return resp.Choices[0].Message.Content, nil
}

func Reset() {
	req.Messages = req.Messages[:1]
}

func GetMessages() []map[string]string {
	var messages []map[string]string
	for i, message := range req.Messages {
		if i == 0 && message.Role == openai.ChatMessageRoleSystem {
			continue
		}
		messages = append(messages, map[string]string{
			"sender":  string(message.Role),
			"content": message.Content,
		})
	}
	return messages
}
