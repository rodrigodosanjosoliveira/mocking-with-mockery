package services

import (
	"errors"

	"github.com/rodrigodosanjosoliveira/mocking-with-mockery/internal/models"
	"github.com/rodrigodosanjosoliveira/mocking-with-mockery/internal/repositories"
)

type ProductService struct {
	repo repositories.ProductRepositoryInterface
}

func NewProductService(
	repo repositories.ProductRepositoryInterface) ProductService {

	return ProductService{
		repo: repo,
	}
}

func (s ProductService) Insert(
	productID string,
	product models.InsertProductDto) error {
	if len(productID) == 0 {
		return errors.New("productID can be not null")
	}

	err := s.repo.Add(models.Product{
		ID:    productID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	})
	if err != nil {
		return err
	}
	return nil
}
