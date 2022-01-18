package v1

import (
	"time"

	"github.com/labstack/echo"
	"github.com/trykafito/kafito/internal/user"
	"github.com/trykafito/kafito/pkg/jwt"
)

type authForm struct {
	Phone    string `json:"phone" form:"phone"`
	Region   string `json:"region" form:"region"`
	Password string `json:"password" form:"password"`
}

func register(ctx echo.Context) error {
	form := new(authForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(400, echo.Map{"error": err.Error()})
	}

	phone, err := user.ParsePhoneNumber(form.Phone, form.Region)
	if err != nil {
		return ctx.JSON(400, echo.Map{"error": err.Error()})
	}

	u := &user.User{
		Phone:    phone,
		Password: form.Password,
	}

	if err := u.Save(); err != nil {
		return ctx.JSON(500, echo.Map{"error": err.Error()})
	}

	t, err := jwt.Create(SecretKey, u.ID.Hex(), time.Now().Add(expireDuration))
	if err != nil {
		return ctx.JSON(500, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(200, echo.Map{
		"message": "successfully registered",
		"token":   t,
		"user":    u,
	})
}
