package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Service interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Logout(c *gin.Context)
	Refresh(c *gin.Context)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	if repo == nil {
		fmt.Fprintln(os.Stderr, "auth service: repository is nil")
		os.Exit(1)
	}

	return &service{repo}
}

func (s *service) Login(c *gin.Context) {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	if req.Username == "" || req.Password == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "username and password are required",
		})
		return
	}

	uid, err := s.repo.Login(req.Username, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "login",
		"uid":     uid,
	})
}

func (s *service) Register(c *gin.Context) {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	if req.Username == "" || req.Password == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "username and password are required",
		})
		return
	}

	uid, err := s.repo.Register(req.Username, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "register",
		"uid":     uid,
	})
}

func (s *service) Logout(c *gin.Context) {
	// TODO: implement
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "unimplemented",
	})
}

func (s *service) Refresh(c *gin.Context) {
	// TODO: implement
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "unimplemented",
	})
}
