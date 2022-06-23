package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"strconv"
)

func ChatBotWebSocketHandler(ws *websocket.Conn) {
	defer ws.Close()

	var keyStr string
	var msg string

	for {
		// Read
		err := websocket.Message.Receive(ws, &keyStr)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			//c.Logger().Error(err)
			return
		}
		key, err := asciiNum(keyStr)

		if isControl(key) {
			switch key {
			case 8:
				if len(msg) > 1 {
					msg = msg[:len(msg)-2]
				} else {
					msg = ""
				}

			case 13:
				fmt.Printf("message: %s\n", msg)
				message = msg
				msg = ""

			case 37, 38, 39, 40:
				fmt.Println("arrow key")
			}
		} else {
			msg += string(key)
		}

		// Write
		if response != "" {
			err = websocket.Message.Send(ws, response)
			if err != nil {
				fmt.Printf("%s\n", err)
			}

			response = ""
		}

	}
}

func asciiNum(str string) (int, error) {
	ascii, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}
	return ascii, nil
}

func isControl(key int) bool {
	return key == 8 || // backspace
		key == 13 || // enter
		key == 37 || // arrow left
		key == 38 || // arrow up
		key == 39 || // arrow right
		key == 40 // arrow down
}
