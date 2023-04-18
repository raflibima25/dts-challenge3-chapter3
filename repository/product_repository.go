package repository

import (
	"challenge-3-chapter-3/entity"
)

type ProductRepository interface {
	FindById(id uint) *entity.Product
	FindAll() []*entity.Product
}
