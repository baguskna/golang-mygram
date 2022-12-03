package repository

import (
	"golang-mygram/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	Delete(ID int) error
	FindAll() ([]domain.User, error)
	FindById(ID int) (domain.User, error)
	Update(user domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) Delete(ID int) error {
	err := r.db.Delete(&domain.User{}, ID).Error
	return err
}

func (r *userRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindById(ID int) (domain.User, error) {
	var user domain.User
	err := r.db.Find(&user, ID).Error
	return user, err
}

func (r *userRepository) Update(user domain.User) (domain.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}
