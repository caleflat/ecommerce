package product

import "github.com/gin-gonic/gin"

func Configure(r *gin.RouterGroup) {
	svc := NewService(nil)
	r.GET("/", svc.GetProducts)
	r.GET("/:id", svc.GetProduct)
	r.POST("/", svc.CreateProduct)
	r.PUT("/:id", svc.UpdateProduct)
	r.DELETE("/:id", svc.DeleteProduct)
}

type Product struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       int64     `json:"price" db:"price"`
	Variants    []Variant `json:"variants" db:"-"`
}

type Variant struct {
	ID          int64   `json:"id" db:"id"`
	ProductID   int64   `json:"product_id" db:"product_id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       int64   `json:"price" db:"price"`
	Images      []Image `json:"images" db:"-"`
}

type Image struct {
	ID        int64  `json:"id" db:"id"`
	VariantID int64  `json:"variant_id" db:"variant_id"`
	URL       string `json:"url" db:"url"`
}
