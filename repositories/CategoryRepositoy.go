package repositories

import (
	"discoverco.co/server/configs"
	"discoverco.co/server/models"
)

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
	FindByID(id uint) (models.Category, error)
	Create(category models.Category) (models.Category, error)
	Update(category models.Category) (models.Category, error)
	Delete(category models.Category) error
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	result := configs.DB.Find(&categories)
	return categories, result.Error
}

func (r *categoryRepository) FindByID(id uint) (models.Category, error) {
	var category models.Category
	result := configs.DB.First(&category, id)
	return category, result.Error
}

func (r *categoryRepository) Create(category models.Category) (models.Category, error) {
	result := configs.DB.Create(&category)
	return category, result.Error
}

func (r *categoryRepository) Update(category models.Category) (models.Category, error) {
	result := configs.DB.Save(&category)
	return category, result.Error
}

func (r *categoryRepository) Delete(category models.Category) error {
	result := configs.DB.Delete(&category)
	return result.Error
}
