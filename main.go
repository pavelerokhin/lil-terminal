package main

/*
	PROTOTYPE!
*/

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"os"
	"os/exec"
	"text/template"

	"github.com/labstack/echo/v4"
)

var (
	stdin  *io.WriteCloser
	stdout *io.ReadCloser

	message, response string // global state variables, see also in server.go
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	var err error

	t := &Template{
		templates: template.Must(template.ParseGlob("client/*.html")),
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = t

	e.Static("/", "./client")

	e.GET("/", index)
	e.GET("/ws", ws)

	stdin, stdout, err = goChatBot()
	defer (*stdout).Close()
	defer (*stdin).Close()

	if err != nil {
		fmt.Printf("error launching chat bot %s", err)
		os.Exit(1)
	}

	// connection to bot
	go func() {
		for {
			if message != "" {

				// input massages to bot
				io.WriteString(*stdin, message)
				if err != nil {
					fmt.Println(err)
					break
				}

				// output massages to bot: bot's reaction
				tmp := make([]byte, 1024)
				_, err = (*stdout).Read(tmp)
				response = string(tmp)
				fmt.Println(response)

				message = ""
			}
		}
	}()

	// start server (NB ws handler)
	e.Logger.Fatal(e.Start(":1323"))
}

func goChatBot() (*io.WriteCloser, *io.ReadCloser, error) {
	cmd := exec.Command("./chatbot/chatbot")
	in, _ := cmd.StdinPipe()
	out, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return &in, &out, nil
}
