package chatbot

import (
	"fmt"
	"strings"
)

/*
	mocks chatbot
*/

type Bot interface {
	StartConversation() string
	ContinueConversation() string
	FinishConversation() string
}

type chatBot struct{}

func GetChatBot() *chatBot {
	return &chatBot{}
}

func (b *chatBot) StartConversation() string {
	return "Hello!"
}

func (b *chatBot) ContinueConversation(message string) string {
	return fmt.Sprintf("OK, you mean: %s", strings.ToLower(message))
}

func (b *chatBot) FinishConversation() string {
	return "See you soon!"
}
