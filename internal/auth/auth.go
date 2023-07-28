package auth

import (
	"github.com/gin-gonic/gin"
)

func Execute() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.SetTrustedProxies(nil)

	svc := NewService(nil)
	r.POST("/login", svc.Login)
	r.POST("/register", svc.Register)
	r.POST("/logout", svc.Logout)
	r.POST("/refresh", svc.Refresh)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
