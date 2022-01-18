package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start(port, secretKey string) error {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.CORS(), middleware.Recover())

	return e.Start(port)
}
