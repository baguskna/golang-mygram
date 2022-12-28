package service

import (
	"golang-mygram/model/domain"
	"golang-mygram/repository"
)

type UserService interface {
	Create(user domain.User) (domain.User, error)
	Delete(ID int) error
	FindAll() ([]domain.User, error)
	FindById(ID int) (domain.User, error)
	Update(ID int, newUser domain.UserUpdate) (domain.User, error)
	Login(user domain.User) (domain.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) *userService {
	return &userService{r}
}

func (s *userService) Create(user domain.User) (domain.User, error) {
	return s.userRepository.Create(user)
}

func (s *userService) Delete(ID int) error {
	return s.userRepository.Delete(ID)
}

func (s *userService) FindAll() ([]domain.User, error) {
	return s.userRepository.FindAll()
}

func (s *userService) FindById(ID int) (domain.User, error) {
	return s.userRepository.FindById(ID)
}

func (s *userService) Update(ID int, newUser domain.UserUpdate) (domain.User, error) {
	user, err := s.userRepository.FindById(ID)

	user.Username = newUser.Username
	user.Email = newUser.Email
	user.Age = newUser.Age

	newestUser, err := s.userRepository.Update(user)

	return newestUser, err
}

func (s *userService) Login(user domain.User) (domain.User, error) {
	return s.userRepository.Login(user)
}
