package v1

import (
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	SecretKey      []byte
	expireDuration = time.Hour * 336
)

func Register(e *echo.Echo, sk string) {
	SecretKey = []byte(sk)

	v1 := e.Group("/v1")

	authGroup := v1.Group("/auth")
	authGroup.POST("/register", register)
	authGroup.POST("/login", login)

	v1.Group("/", middleware.JWT(SecretKey), setUser)
}
