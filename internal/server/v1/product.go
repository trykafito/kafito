package v1

import (
	"github.com/labstack/echo"
	"github.com/trykafito/kafito/internal/product"
	"github.com/trykafito/kafito/internal/user"
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
