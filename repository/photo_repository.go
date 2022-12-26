package repository

import (
	"golang-mygram/model/domain"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo domain.Photo) (domain.Photo, error)
	Delete(ID int) error
	FindAll() ([]domain.Photo, error)
	FindById(ID int) (domain.Photo, error)
	Update(photo domain.Photo) (domain.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) Create(photo domain.Photo) (domain.Photo, error) {
	err := r.db.Create(&photo).Error
	return photo, err
}

func (r *photoRepository) Delete(ID int) error {
	err := r.db.Delete(&domain.Photo{}, ID).Error
	return err
}

func (r *photoRepository) FindAll() ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.Preload("User").Find(&photos).Error
	return photos, err
}

func (r *photoRepository) FindById(ID int) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Preload("User").Find(&photo, ID).Error
	return photo, err
}

func (r *photoRepository) Update(photo domain.Photo) (domain.Photo, error) {
	err := r.db.Save(&photo).Error
	return photo, err
}
