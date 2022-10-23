package http

import (
	"go-api/model"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/product", w.getAll)
	e.POST("/product", w.createProduct)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, w.Products)
}

func (w WebServer) createProduct(c echo.Context) error {
	product := model.Product{
		ID: uuid.NewV4().String(),
	}

	if err := c.Bind(product); err != nil {
		return nil
	}
	w.Products.Add(product)
	return c.JSON(http.StatusCreated, product)
}
