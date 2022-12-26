package repository

import (
	"golang-mygram/model/domain"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment domain.Comment) (domain.Comment, error)
	Delete(ID int) error
	FindAll() ([]domain.Comment, error)
	FindById(ID int) (domain.Comment, error)
	Update(comment domain.Comment) (domain.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Create(comment domain.Comment) (domain.Comment, error) {
	err := r.db.Create(&comment).Error
	return comment, err
}

func (r *commentRepository) Delete(ID int) error {
	err := r.db.Delete(&domain.Comment{}, ID).Error
	return err
}

func (r *commentRepository) FindAll() ([]domain.Comment, error) {
	var comments []domain.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comments).Error
	return comments, err
}

func (r *commentRepository) FindById(ID int) (domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comment, ID).Error
	return comment, err
}

func (r *commentRepository) Update(comment domain.Comment) (domain.Comment, error) {
	err := r.db.Save(&comment).Error
	return comment, err
}
