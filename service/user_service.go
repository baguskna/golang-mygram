package service

import (
	"golang-mygram/model/domain"
	"golang-mygram/repository"
	"golang-mygram/utils"
)

type UserService interface {
	Create(user domain.User) (domain.UserResponse, error)
	Delete(ID int) error
	FindAll() ([]domain.UserResponse, error)
	FindById(ID int) (domain.UserResponse, error)
	Update(ID int, newUser domain.UserUpdate) (domain.UserResponse, error)
	Login(user domain.User) (domain.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) *userService {
	return &userService{r}
}

func (s *userService) Create(user domain.User) (domain.UserResponse, error) {
	user, err := s.userRepository.Create(user)
	userResponse := utils.UserResponseFunc(user)
	return userResponse, err
}

func (s *userService) Delete(ID int) error {
	return s.userRepository.Delete(ID)
}

func (s *userService) FindAll() ([]domain.UserResponse, error) {
	users, err := s.userRepository.FindAll()

	var usersResponse []domain.UserResponse
	for _, user := range users {
		userResponse := utils.UserResponseFunc(user)
		usersResponse = append(usersResponse, userResponse)
	}

	return usersResponse, err
}

func (s *userService) FindById(ID int) (domain.UserResponse, error) {
	user, err := s.userRepository.FindById(ID)
	userResponse := utils.UserResponseFunc(user)
	return userResponse, err
}

func (s *userService) Update(ID int, newUser domain.UserUpdate) (domain.UserResponse, error) {
	user, err := s.userRepository.FindById(ID)

	user.Username = newUser.Username
	user.Email = newUser.Email
	user.Age = newUser.Age

	newestUser, err := s.userRepository.Update(user)
	userResponse := utils.UserResponseFunc(newestUser)
	return userResponse, err
}

func (s *userService) Login(user domain.User) (domain.User, error) {
	return s.userRepository.Login(user)
}
