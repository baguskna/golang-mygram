package service

import (
	"golang-mygram/model/domain"
	"golang-mygram/repository"
)

type SocialMediaService interface {
	Create(socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	Delete(ID int) (domain.SocialMedia, error)
	FindAll() ([]domain.SocialMedia, error)
	FindById(ID int) (domain.SocialMedia, error)
	Update(ID int, newSocialMedia domain.SocialMediaUpdate) (domain.SocialMedia, error)
}

type socialMediaService struct {
	socialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(r repository.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{r}
}

func (s *socialMediaService) Create(socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	return s.socialMediaRepository.Create(socialMedia)
}

func (s *socialMediaService) Delete(ID int) (domain.SocialMedia, error) {
	socialMedia, err := s.socialMediaRepository.FindById(ID)
	err = s.socialMediaRepository.Delete(ID)
	return socialMedia, err
}

func (s *socialMediaService) FindAll() ([]domain.SocialMedia, error) {
	return s.socialMediaRepository.FindAll()
}

func (s *socialMediaService) FindById(ID int) (domain.SocialMedia, error) {
	return s.socialMediaRepository.FindById(ID)
}

func (s *socialMediaService) Update(ID int, newSocialMedia domain.SocialMediaUpdate) (domain.SocialMedia, error) {
	socialMedia, err := s.socialMediaRepository.FindById(ID)

	socialMedia.Name = newSocialMedia.Name
	socialMedia.SocialMediaURL = newSocialMedia.SocialMediaURL

	newestSocialMedia, err := s.socialMediaRepository.Update(socialMedia)

	return newestSocialMedia, err
}
