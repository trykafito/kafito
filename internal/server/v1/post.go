package v1

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/trykafito/kafito/internal/post"
	"github.com/trykafito/kafito/internal/user"
	"go.mongodb.org/mongo-driver/bson"
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

func editPost(ctx echo.Context) error {
	filter := bson.M{}

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(404, echo.Map{"error": err.Error()})
	}

	filter["_id"] = id

	form := new(postForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(400, echo.Map{"error": err.Error()})
	}

	p, err := post.FindOne(filter)
	if err != nil {
		return ctx.JSON(404, echo.Map{"error": err.Error()})
	}

	p.Title = form.Title
	p.Description = form.Description
	p.Tags = form.Tags
	p.Thumbnail = form.Thumbnail

	if err := p.Save(); err != nil {
		return ctx.JSON(500, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(200, echo.Map{
		"message": "post updated successfully",
		"post":    postToJSON(*p),
	})
}

func getPost(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(400, echo.Map{"error": err.Error()})
	}

	p, err := post.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(404, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(200, echo.Map{
		"post": postToJSON(*p),
	})
}

func listPosts(ctx echo.Context) error {
	filter := bson.M{}

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	if q := ctx.QueryParam("q"); q != "" {
		filter["title"] = bson.M{"$regex": q, "$options": "i"}
	}

	posts, err := post.Find(filter, page, limit, ctx.Get("sort").(bson.D)...)
	if err != nil {
		return ctx.JSON(404, echo.Map{"error": err.Error()})
	}

	count := post.Count(filter)

	result := []M{}
	for _, p := range posts {
		result = append(result, postToJSON(p))
	}

	return ctx.JSON(200, echo.Map{
		"posts":       result,
		"page":        page,
		"total_count": count,
	})
}
