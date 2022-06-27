package main

import (
	"fmt"
	"strings"
)

/*
	chatbot_conversational
*/

type Bot interface {
	StartConversation() string
	ContinueConversation() string
	FinishConversation() string
	MockCharm() string
}

type ChatBot struct{}

func GetChatBot() *ChatBot {
	return &ChatBot{}
}

func (b *ChatBot) StartConversation() string {
	return "Hello!"
}

func (b *ChatBot) ContinueConversation(message string) string {
	return fmt.Sprintf("OK, you mean: %s", strings.ToLower(message))
}

func (b *ChatBot) FinishConversation() string {
	return "See you soon!"
}

func (b *ChatBot) MockCharm() string {
	return "See you soon!"
}
