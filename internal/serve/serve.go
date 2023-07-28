package serve

import (
	"github.com/caleflat/ecommerce/internal/auth"
	"github.com/caleflat/ecommerce/internal/product"
	"github.com/caleflat/ecommerce/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())

	product.Configure(router.Group("/products"))
	auth.Configure(router.Group("/auth"))
}
