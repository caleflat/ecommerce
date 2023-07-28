package product

import "github.com/jmoiron/sqlx"

type Repository interface {
	GetProducts() ([]Product, error)
	GetProduct(id int64) (Product, error)
	CreateProduct(product Product) (int64, error)
	UpdateProduct(product Product) error
	DeleteProduct(id int64) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetProducts() ([]Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetProduct(id int64) (Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) CreateProduct(product Product) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) UpdateProduct(product Product) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) DeleteProduct(id int64) error {
	//TODO implement me
	panic("implement me")
}
