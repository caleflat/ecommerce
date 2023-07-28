package auth

import (
	"github.com/gin-gonic/gin"
)

func Configure(r *gin.RouterGroup) {
	svc := NewService(nil)
	r.POST("/login", svc.Login)
	r.POST("/register", svc.Register)
	r.POST("/logout", svc.Logout)
	r.POST("/refresh", svc.Refresh)
}
