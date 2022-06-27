package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CMD_STOP = "STOP"
)

var (
	cb = GetChatBot()
)

func main() {
	var command, message, response string

	scanner := bufio.NewScanner(os.Stdin)

conv:
	for scanner.Scan() {
		message = scanner.Text()
		// convert CRLF to LF
		message = strings.Replace(message, "\n", "", -1)

		response, command = answer(message)

		fmt.Fprint(os.Stdout, response)

		switch command {
		case CMD_STOP:
			command = ""
			break conv
		}
	}
}

func answer(message string) (string, string) {
	var response, command string

	switch message {
	case "START":
		response = cb.StartConversation()
	case "END":
		response = cb.FinishConversation()
		command = CMD_STOP
	default:
		response = cb.ContinueConversation(message)
	}

	return response + "\n", command
}
