package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	v1 "github.com/trykafito/kafito/internal/server/v1"
)

func Start(port, secretKey string) error {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.CORS(), middleware.Recover())

	v1.Register(e, secretKey)

	return e.Start(port)
}
