package service

import (
	"ewallet-blockhain/modules/v1/utilities/product/models"
	"ewallet-blockhain/modules/v1/utilities/product/repository"
)

type Service interface {
	ListProduct() ([]models.Product, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) ListProduct() ([]models.Product, error) {
	allproduct, err := s.repository.ListProduct()
	if err != nil {
		return nil, err
	}
	return allproduct, nil
}
