package web

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Start() func(ctx context.Context) error {
	e := echo.New()
	e.HideBanner = true
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()

	return e.Shutdown
}
