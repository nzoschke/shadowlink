package web

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nzoschke/shadowlink"
)

func Start() func(ctx context.Context) error {
	e := echo.New()
	e.HideBanner = true
	e.GET("/*", echo.WrapHandler(http.FileServer(shadowlink.Build())))

	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()

	return e.Shutdown
}
