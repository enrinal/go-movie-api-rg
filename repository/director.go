package repository

import (
	"go-movie-api/model"
	"gorm.io/gorm"
)

type DirectorRepository interface {
	FetchAll() ([]model.Director, error)
	FetchByID(id int) (*model.Director, error)
	Store(s *model.Director) error
}

type directorRepoImpl struct {
	db *gorm.DB
}

func NewDirectorRepo(db *gorm.DB) *directorRepoImpl {
	return &directorRepoImpl{db}
}

func (s *directorRepoImpl) FetchAll() ([]model.Director, error) {
	var directors []model.Director
	result := s.db.Find(&directors)
	if result.Error != nil {
		return nil, result.Error
	}

	return directors, nil
}

func (s *directorRepoImpl) FetchByID(id int) (*model.Director, error) {
	var director model.Director
	result := s.db.First(&director, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &director, nil
}

func (s *directorRepoImpl) Store(director *model.Director) error {
	result := s.db.Create(director)
	if result.Error != nil {
		return result.Error
	}

	return nil
}