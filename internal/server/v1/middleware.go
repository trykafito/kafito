package v1

import (
	"github.com/labstack/echo"
	"github.com/trykafito/kafito/internal/user"
)

func setUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		u, err := user.LoadByRequest(SecretKey, ctx.Request())
		if err != nil {
			return ctx.JSON(401, echo.Map{"error": err.Error()})
		}

		ctx.Set("user", u)
		return next(ctx)
	}
}
