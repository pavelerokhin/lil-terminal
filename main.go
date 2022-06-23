package main

/*
	PROTOTYPE!
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

var (
	clientMessage = make(chan string)
	botResponse   = make(chan string)

	ws *websocket.Conn
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("client/*.html")),
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = t

	e.Static("/", "./client")

	e.GET("/", index)
	e.GET("/ws", webSocketHandler)

	// keep connection to bot
	go backEndToBotConn()
	// keep connection to frontend
	go backEndToFrontEndConn()

	// start server (NB web socket handler)
	e.Logger.Fatal(e.Start(":1323"))
}

// TODO: make it available from everywhere
type message struct {
	Message string `json:"message"`
}

func backEndToBotConn() {
	fmt.Println("start the connection between backend and chat-bot")

	for {
		m := <-clientMessage
		fmt.Println("client message:", m)

		msg := message{
			Message: m,
		}

		json_data, err := json.Marshal(msg)

		if err != nil {
			fmt.Printf("error: %s\n", err)
			break
		}

		resp, err := http.Post("http://localhost:7234/", "application/json", bytes.NewBuffer(json_data))
		if err != nil {
			fmt.Printf("error: %s\n", err)
			break
		}

		// output massages from bot: the reaction
		r, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			break
		}
		botResponse <- string(r)
	}
	fmt.Println("connection between backend and chat-bot has been interrupted")
	os.Exit(1)
}

func backEndToFrontEndConn() {
	fmt.Println("start the connection between backend and frontend")
	for {
		r := <-botResponse
		fmt.Println("bot response:", r)

		err := websocket.Message.Send(ws, r)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			break
		}
	}
	fmt.Println("connection between backend and frontend has been interrupted")
	os.Exit(1)
}

// handlers
func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "")
}

func webSocketHandler(c echo.Context) error {
	websocket.Handler(ChatBotWebSocketHandler).ServeHTTP(c.Response(), c.Request())
	return nil
}
