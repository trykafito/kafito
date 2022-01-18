package v1

import (
	"time"

	"github.com/labstack/echo"
)

var (
	SecretKey      []byte
	expireDuration = time.Hour * 336
)

func Register(e *echo.Echo, secretKey string) {
	SecretKey = []byte(secretKey)

	v1 := e.Group("/v1")

	authGroup := v1.Group("/auth")
	authGroup.POST("/register", register)
	authGroup.POST("/login", login)
}
