package repository

import (
	"golang-mygram/model/domain"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	Delete(ID int) error
	FindAll() ([]domain.SocialMedia, error)
	FindById(ID int) (domain.SocialMedia, error)
	Update(socialMedia domain.SocialMedia) (domain.SocialMedia, error)
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) Create(socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	err := r.db.Create(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) Delete(ID int) error {
	return r.db.Delete(&domain.SocialMedia{}, ID).Error
}

func (r *socialMediaRepository) FindAll() ([]domain.SocialMedia, error) {
	var socialMedias []domain.SocialMedia
	err := r.db.Preload("User").Find(&socialMedias).Error
	return socialMedias, err
}

func (r *socialMediaRepository) FindById(ID int) (domain.SocialMedia, error) {
	var socialMedia domain.SocialMedia
	err := r.db.Preload("User").Find(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) Update(socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	err := r.db.Save(&socialMedia).Error
	return socialMedia, err
}
