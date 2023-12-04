package application

import (
	"database/sql"
	"errors"
	"golang_track_expense/domain/entity"
	"golang_track_expense/domain/repository"

	"github.com/google/uuid"
)

var (
	ErrUserAlreadyExist = errors.New("user already exists")
	ErrUserNotFound     = errors.New("user not found")
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (u *UserService) List() ([]entity.User, error) {
	return u.UserRepository.List()
}

func (u *UserService) GetByUserId(id string) (entity.User, error) {
	var user entity.User

	uid, err := uuid.Parse(id)
	if errors.Is(err, sql.ErrNoRows) {
		return user, ErrUserNotFound
	}

	return u.UserRepository.GetByUserId(uid)
}

func (u *UserService) Create(name string) error {

	_, err := u.GetByUserName(name)
	if err == nil {
		return ErrUserAlreadyExist
	}

	user, err := entity.NewUser(name)
	if err != nil {
		return err
	}

	return u.UserRepository.Create(user)
}

func (u *UserService) GetByUserName(username string) (entity.User, error) {
	return u.UserRepository.GetByUserName(username)
}
