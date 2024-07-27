package repositories

import (
	"discoverco.co/server/configs"
	"discoverco.co/server/models"
)

type FeatureRepository interface {
	FindAll() ([]models.Feature, error)
}

type FeatureRepositoryImpl struct{}

func NewFeatureRepositoryImpl() FeatureRepository {
	return &FeatureRepositoryImpl{}
}

// FindAll implements FeatureRepository.
func (f *FeatureRepositoryImpl) FindAll() ([]models.Feature, error) {
	var features []models.Feature
	result := configs.DB.Find(&features)
	return features, result.Error
}
