package services

import (
	"discoverco.co/server/models"
	"discoverco.co/server/repositories"
)

type CategoryService interface {
	FindAll() ([]models.Category, error)
	FindByID(id uint) (models.Category, error)
	Create(category models.Category) (models.Category, error)
	Update(id uint, input models.Category) (models.Category, error)
	Delete(id uint) error
}

type categoryService struct {
	repository repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repository: repo}
}

func (s *categoryService) FindAll() ([]models.Category, error) {
	return s.repository.FindAll()
}

func (s *categoryService) FindByID(id uint) (models.Category, error) {
	return s.repository.FindByID(id)
}

func (s *categoryService) Create(category models.Category) (models.Category, error) {
	return s.repository.Create(category)
}

func (s *categoryService) Update(id uint, input models.Category) (models.Category, error) {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return models.Category{}, err
	}

	category.Name = input.Name
	return s.repository.Update(category)
}

func (s *categoryService) Delete(id uint) error {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}
	return s.repository.Delete(category)
}
