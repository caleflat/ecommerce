package product

import "github.com/gin-gonic/gin"

type Service interface {
	GetProducts(c *gin.Context)
	GetProduct(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetProducts(c *gin.Context) {
	products, err := s.repo.GetProducts()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"products": products,
	})
}

func (s *service) GetProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *service) CreateProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
