package main

import (
	"errors"
	"log"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func serveHome(c echo.Context) error {
	r := c.Request()
	w := c.Response()
	log.Print(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return nil
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed please send a GET request", http.StatusMethodNotAllowed)
		return nil
	}
	http.ServeFile(w, r, "index.html")
	return nil
}

func main() {
	hub := newHub()
	go hub.run()
	e := echo.New()
	e.GET("/", serveHome)
	e.GET("/ws", func(c echo.Context) error {
		serveWs(hub, c)
		return nil
	})
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
