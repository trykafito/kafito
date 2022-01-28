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

type M map[string]interface{}

func Register(e *echo.Echo, sk string) {
	SecretKey = []byte(sk)

	v1 := e.Group("/v1")

	authGroup := v1.Group("/auth")
	authGroup.POST("/register", register)
	authGroup.POST("/login", login)

	r := v1.Group("/", middleware.JWT(SecretKey), setUser)

	currentUserGroup := r.Group("current-user")
	currentUserGroup.GET("", getCurrentUser)

	productGroup := r.Group("products")
	productGroup.POST("", addProduct)
	productGroup.PUT("/:id", editProduct)

	postGroup := r.Group("posts")
	postGroup.POST("", addPost)
}
