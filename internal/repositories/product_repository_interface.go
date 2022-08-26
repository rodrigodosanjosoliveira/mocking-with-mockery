package repositories

import "github.com/rodrigodosanjosoliveira/mocking-with-mockery/internal/models"

type ProductRepositoryInterface interface {
	Add(product models.Product) error
}
