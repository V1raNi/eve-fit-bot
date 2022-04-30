package api

import (
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app, db *gorm.DB) {
	//h := &handler{
	//	DB: db,
	//}

	//routes := app.Group("/fits")
	//routes.Get("/")
	//routes.Post("/", h.AddProduct)
	//routes.Get("/", h.GetProducts)
	//routes.Get("/:id", h.GetProduct)
	//routes.Put("/:id", h.UpdateProduct)
	//routes.Delete("/:id", h.DeleteProduct)
}
