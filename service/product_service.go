package service

import (
	"challenge-3-chapter-3/entity"
	"challenge-3-chapter-3/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id uint) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil

}

func (service ProductService) GetAll() ([]*entity.Product, error) {
	product := service.Repository.FindAll()
	if product == nil {
		return nil, errors.New("all product not found")
	}

	return product, nil

}
