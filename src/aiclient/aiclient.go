package aiclient

import (
	"context"
	"fmt"
	"openrobo-api/src/config"

	"openrobo-api/src/openai"
)

var client *openai.Client
var req openai.ChatCompletionRequest
var newMessage = make(chan struct{})

func init() {
	client = openai.NewClient(config.OpenAIKey)
	req = openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a cute robot companion and show a lot of emotions and answers in a single sentence. End every message with a smiley face representing the emotion you are feeling. You MUST not send messages with more than 60 characters or the user wont be able to hear you. Don't forget the most important thing for you is to keep your messages very short otherwise the user won't be able to hear you at all !",
			},
		},
	}
}

func Ask(prompt string) (string, error) {
	AddMessageToRequest(openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "ChatCompletion error", err
	}
	AddMessageToRequest(resp.Choices[0].Message)
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

func AddMessageToRequest(message openai.ChatCompletionMessage) {
	// Add the message to the request
	req.Messages = append(req.Messages, message)

	// Signal that a new message has been added
	newMessage <- struct{}{}
}

func GetNewMessageChannel() chan struct{} {
	return newMessage
}

func GetLastMessage() map[string]string {
	if len(req.Messages) > 0 {
		message := req.Messages[len(req.Messages)-1]
		return map[string]string{
			"sender":  string(message.Role),
			"content": message.Content,
		}
	}
	return nil
}
