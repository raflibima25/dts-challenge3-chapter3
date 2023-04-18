package service

import (
	"challenge-3-chapter-3/entity"
	"challenge-3-chapter-3/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {

	productRepository.Mock.On("FindById", uint(1)).Return(nil)

	product, err := productService.GetOneProduct(1)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error())

}

func TestGetOneProductFound(t *testing.T) {
	product := entity.Product{
		ID:          uint(2),
		Title:       "Harry Potter",
		Description: "J.K Rowling",
	}

	productRepository.Mock.On("FindById", uint(2)).Return(product)

	result, err := productService.GetOneProduct(2)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.ID, result.ID)
	assert.Equal(t, product.Title, result.Title)
}

func TestGetAllProductNotFound(t *testing.T) {

	productRepository.Mock.On("FindAll").Return(nil)

	result, err := productService.GetAll()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "all product not found", err.Error())
}

func TestGetAllProductFound(t *testing.T) {

	productRepository := &repository.ProductRepositoryMock{}
	productService := ProductService{Repository: productRepository}

	product := []*entity.Product{
		{ID: uint(4), Title: "Harry Potter", Description: "J.K Rowling"},
		{ID: uint(5), Title: "The Lion", Description: "C.S Lewis"},
		{ID: uint(6), Title: "Pinocchio", Description: "Carlo Collodi"},
	}

	productRepository.Mock.On("FindAll").Return(product)

	res, err := productService.GetAll()

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, product, res)
}
