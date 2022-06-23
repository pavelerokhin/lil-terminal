package main

import (
	"fmt"
	"io"
	"os/exec"
)

var (
	cb = GetChatBot()
)

//func main() {
//	var message string
//
//	fmt.Println(cb.StartConversation())
//
//	for {
//		n, err := fmt.Scan(&message)
//		if n > 0 && err != nil {
//			fmt.Printf("error getting message, %s", err)
//			break
//		}
//
//		if message == "c" {
//			break
//		}
//
//		fmt.Println(cb.ContinueConversation(message))
//	}
//	fmt.Println(cb.FinishConversation())
//}

func main() {
	cmd := exec.Command("./chatbot")

	stdout, _ := cmd.StdoutPipe()
	defer stdout.Close()

	stdin, _ := cmd.StdinPipe()
	defer stdin.Close()

	io.WriteString(stdin, "an old falcon")

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))

		io.WriteString(stdin, "an")
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
