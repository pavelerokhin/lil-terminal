package main

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"net/http"
)

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "")
}

func ws(c echo.Context) error {
	websocket.Handler(ChatBotWebSocketHandler).ServeHTTP(c.Response(), c.Request())
	return nil
}
