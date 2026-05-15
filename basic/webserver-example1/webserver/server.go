package webserver

import (
	"github.com.br/tonnytg/std-golang-wiki/live-go-esquenta-fc/data"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WebServer struct {
}

func NewWebserver() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/product", w.getAll)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, data.Products)
}