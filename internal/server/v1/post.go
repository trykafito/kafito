package v1

import (
	"github.com/labstack/echo"
	"github.com/trykafito/kafito/internal/post"
	"github.com/trykafito/kafito/internal/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type postForm struct {
	Title       string             `json:"title" form:"title"`
	Description string             `json:"description" form:"description"`
	Tags        []string           `json:"tags" form:"tags"`
	Thumbnail   primitive.ObjectID `json:"thumbnail" form:"thumbnail"`
}

func postToJSON(p post.Post) M {
	return M{
		"id":          p.ID.Hex(),
		"created_by":  p.CreatedBy,
		"title":       p.Title,
		"description": p.Description,
		"tags":        p.Tags,
		"thumbnail":   p.Thumbnail,
		"created_at":  p.CreatedAt,
	}
}

func addPost(ctx echo.Context) error {
	u := ctx.Get("user").(*user.User)

	form := new(postForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(400, echo.Map{"error": err.Error()})
	}

	p := &post.Post{
		CreatedBy:   u.ID,
		Title:       form.Title,
		Description: form.Description,
		Tags:        form.Tags,
		Thumbnail:   form.Thumbnail,
	}

	if err := p.Save(); err != nil {
		return ctx.JSON(500, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(200, echo.Map{
		"message": "post created successfully",
		"post":    postToJSON(*p),
	})
}
