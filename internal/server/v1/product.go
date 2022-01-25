package v1

import (
	"github.com/labstack/echo"
	"github.com/trykafito/kafito/internal/product"
	"github.com/trykafito/kafito/internal/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productForm struct {
	Title        string                 `json:"title" form:"title"`
	Description  string                 `json:"description" form:"description"`
	Informations map[string]interface{} `json:"informations" form:"informations"`
	Price        float64                `json:"price" form:"price"`
	Quantity     int                    `json:"quantity" form:"quantity"`
	Thumbnail    primitive.ObjectID     `json:"thumbnail" form:"thumbnail"`
}

func productToJSON(p product.Product) M {
	return M{
		"id":           p.ID.Hex(),
		"created_by":   p.CreatedBy,
		"title":        p.Title,
		"description":  p.Description,
		"informations": p.Informations,
		"price":        p.Price,
		"quantity":     p.Quantity,
		"thumbnail":    p.Thumbnail,
		"created_at":   p.CreatedAt,
	}
}

func addProduct(ctx echo.Context) error {
	u := ctx.Get("user").(*user.User)

	form := new(productForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(400, echo.Map{"error": err.Error()})
	}

	p := &product.Product{
		CreatedBy:    u.ID,
		Title:        form.Title,
		Description:  form.Description,
		Informations: form.Informations,
		Price:        form.Price,
		Quantity:     form.Quantity,
		Thumbnail:    form.Thumbnail,
	}

	if err := p.Save(); err != nil {
		return ctx.JSON(500, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(200, echo.Map{
		"message": "product created successfully",
		"product": productToJSON(*p),
	})
}

func editProduct(ctx echo.Context) error {
	filter := bson.M{}

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(404, echo.Map{"error": err.Error()})
	}

	filter["_id"] = id

	form := new(productForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(400, echo.Map{"error": err.Error()})
	}

	p, err := product.FindOne(filter)
	if err != nil {
		return ctx.JSON(404, echo.Map{"error": err.Error()})
	}

	p.Title = form.Title
	p.Description = form.Description
	p.Informations = form.Informations
	p.Price = form.Price
	p.Quantity = form.Quantity
	p.Thumbnail = form.Thumbnail

	if err := p.Save(); err != nil {
		return ctx.JSON(500, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(200, echo.Map{
		"message": "product updated successfully",
		"product": productToJSON(*p),
	})
}
