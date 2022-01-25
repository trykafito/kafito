package v1

import (
	"github.com/labstack/echo"
	"github.com/trykafito/kafito/internal/user"
	"go.mongodb.org/mongo-driver/bson"
)

func userToJSON(u user.User) bson.M {
	return bson.M{
		"id":         u.ID.Hex(),
		"name":       u.Name,
		"phone":      u.Phone,
		"password":   u.Password,
		"super_user": u.SuperUser,
		"created_at": u.CreatedAt,
	}
}

func getCurrentUser(ctx echo.Context) error {
	u := ctx.Get("user").(*user.User)

	return ctx.JSON(200, echo.Map{
		"user": userToJSON(*u),
	})
}
