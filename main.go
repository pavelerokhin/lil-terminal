package main

/*
	PROTOTYPE!
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"text/template"
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
	//go backEndToBotConnConversational()
	go backEndToBotConnStdInOut()
	// keep connection to frontend
	go backEndToFrontEndConn()

	// start server (NB web socket handler)
	e.Logger.Fatal(e.Start(":1323"))
}

// TODO: make it available from everywhere
type message struct {
	Message string `json:"message"`
}

func backEndToBotConnConversational() {
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

func backEndToBotConnStdInOut() {
	fmt.Println("start the connection between backend and chat-bot")

	var err error
	//var stdout, stderr bytes.Buffer

	subProcess := exec.Command("./chatbot_std/chatbot_std") //Just for testing, replace with your subProcess
	stdin, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Printf("error getting stdIn pipe: %s", err)
		return
	}

	//subProcess.Stdout = &stdout
	//subProcess.Stderr = &stderr

	//stdout, err := subProcess.StdoutPipe()
	//if err != nil {
	//	fmt.Printf("error getting stdOut pipe: %s", err)
	//	return
	//}
	//
	//stderr, err := subProcess.StderrPipe()
	//if err != nil {
	//	fmt.Printf("error getting stdError pipe: %s", err)
	//	return
	//}
	var stdout, stderr bytes.Buffer
	subProcess.Stdout = &stdout // standard output
	subProcess.Stderr = &stderr // standard error

	err = subProcess.Start()
	if err != nil {
		fmt.Printf("error starting subprocess: %s", err)
		return
	}

	//stdoutScanner := bufio.NewScanner(stdout)
	//stderrScanner := bufio.NewScanner(stderr)
	//go func() {
	//	for stdoutScanner.Scan() {
	//		text := stdoutScanner.Text()
	//		println(text)
	//		botResponse <- text
	//	}
	//}()
	//
	//go func() {
	//	for stderrScanner.Scan() {
	//		text := stderrScanner.Text()
	//		fmt.Println("error in stdError: ", text)
	//		os.Exit(1)
	//	}
	//}()

	//go func() {
	//	for {
	//		if len(stdout.String()) > 0 {
	//			text := stdout.String()
	//			println(text)
	//			botResponse <- text
	//		}
	//	}
	//}()

	for {
		m := <-clientMessage
		fmt.Println("client message:", m)

		_, err1 := io.WriteString(stdin, m)
		if err1 != nil {
			fmt.Println("error writing to stdIn:", err1)
			break
		}

		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

		botResponse <- outStr
	}

	fmt.Println("connection between backend and chat-bot has been interrupted")
	os.Exit(1)
}

func runProcess(process *exec.Cmd) (stdout, stderr string, err error) {
	var stdoutbuf, stderrbuf bytes.Buffer
	process.Stdout = &stdoutbuf
	process.Stderr = &stderrbuf
	if err := process.Run(); err != nil {
		return "", "", err
	}
	return stdoutbuf.String(), stderrbuf.String(), nil
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
