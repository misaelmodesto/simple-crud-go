package usecase

import (
	"github.com/misaelmodesto/simple-crud-go/model"
	"github.com/misaelmodesto/simple-crud-go/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GETProducts()
}
