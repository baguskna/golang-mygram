package service

import (
	"golang-mygram/model/domain"
	"golang-mygram/repository"
)

type CommentService interface {
	Create(comment domain.Comment) (domain.Comment, error)
	Delete(ID int) (domain.Comment, error)
	FindAll() ([]domain.Comment, error)
	FindById(ID int) (domain.Comment, error)
	Update(ID int, newComment domain.CommentUpdate) (domain.Comment, error)
}

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(r repository.CommentRepository) *commentService {
	return &commentService{r}
}

func (s *commentService) Create(comment domain.Comment) (domain.Comment, error) {
	return s.commentRepository.Create(comment)
}

func (s *commentService) Delete(ID int) (domain.Comment, error) {
	comment, err := s.commentRepository.FindById(ID)
	err = s.commentRepository.Delete(ID)
	return comment, err
}

func (s *commentService) FindAll() ([]domain.Comment, error) {
	return s.commentRepository.FindAll()
}

func (s *commentService) FindById(ID int) (domain.Comment, error) {
	return s.commentRepository.FindById(ID)
}

func (s *commentService) Update(ID int, newComment domain.CommentUpdate) (domain.Comment, error) {
	comment, err := s.commentRepository.FindById(ID)

	comment.Message = newComment.Message

	newestComment, err := s.commentRepository.Update(comment)

	return newestComment, err
}
