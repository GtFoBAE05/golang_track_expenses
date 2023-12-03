package application

import (
	"golang_track_expense/domain/entity"
	"golang_track_expense/domain/repository"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

type UserRequest struct {
	Name string `json:"name"`
}

func (u *UserService) List() ([]entity.User, error) {
	return u.UserRepository.List()
}

func (u *UserService) GetByUserId(id string) (entity.User, error) {
	var user entity.User

	uid, err := uuid.Parse(id)
	if err != nil {
		return user, err
	}

	return u.UserRepository.GetByUserId(uid)
}

func (u *UserService) Create(name string) error {
	return u.UserRepository.Create(name)
}

func (u *UserService) GetByUserName(username string) (entity.User, error) {
	return u.UserRepository.GetByUserName(username)
}
