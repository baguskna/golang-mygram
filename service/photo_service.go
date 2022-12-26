package service

import (
	"golang-mygram/model/domain"
	"golang-mygram/repository"
)

type PhotoService interface {
	Create(photo domain.Photo) (domain.Photo, error)
	Delete(ID int) (domain.Photo, error)
	FindAll() ([]domain.Photo, error)
	FindById(ID int) (domain.Photo, error)
	Update(ID int, newPhoto domain.PhotoUpdate) (domain.Photo, error)
}

type photoService struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoService(r repository.PhotoRepository) *photoService {
	return &photoService{r}
}

func (s *photoService) Create(photo domain.Photo) (domain.Photo, error) {
	return s.photoRepository.Create(photo)
}

func (s *photoService) Delete(ID int) (domain.Photo, error) {
	photo, err := s.photoRepository.FindById(ID)
	err = s.photoRepository.Delete(ID)
	return photo, err
}

func (s *photoService) FindAll() ([]domain.Photo, error) {
	return s.photoRepository.FindAll()
}

func (s *photoService) FindById(ID int) (domain.Photo, error) {
	return s.photoRepository.FindById(ID)
}

func (s *photoService) Update(ID int, newPhoto domain.PhotoUpdate) (domain.Photo, error) {
	photo, err := s.photoRepository.FindById(ID)

	photo.Title = newPhoto.Title
	photo.Caption = newPhoto.Caption
	photo.PhotoURL = newPhoto.PhotoURL

	newestPhoto, err := s.photoRepository.Update(photo)

	return newestPhoto, err
}
