package services

import (
	"discoverco.co/server/models"
	"discoverco.co/server/repositories"
)

type FeatureService interface {
	FindAll() ([]models.Feature, error)
}

type FeatureServiceImpl struct {
	repository repositories.FeatureRepository
}

func NewFeatureService(repo repositories.FeatureRepository) FeatureService {
	return &FeatureServiceImpl{repository: repo}
}

// FindAll implements FeatureService.
func (f *FeatureServiceImpl) FindAll() ([]models.Feature, error) {
	return f.repository.FindAll()
}
