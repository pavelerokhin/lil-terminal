package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	cb = GetChatBot()
)

func main() {
	e := echo.New()
	e.POST("/", say) // curl -X POST http://0.0.0.0:1111/ -d 'message=START'
	e.Logger.Fatal(e.Start("0.0.0.0:7234"))
}

// TODO: make it available from everywhere
type message struct {
	Message string `json:"message"`
}

func say(c echo.Context) error {
	var response string
	m := &message{}

	if err := c.Bind(m); err != nil {
		return err
	}

	fmt.Printf("message has arrived: %s\n", m.Message)

	switch m.Message {
	case "START":
		response = cb.StartConversation()
	case "END":
		response = cb.FinishConversation()
	default:
		response = cb.ContinueConversation(m.Message)
	}

	return c.String(http.StatusOK, response)
}
