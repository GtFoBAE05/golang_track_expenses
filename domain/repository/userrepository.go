package repository

import (
	"golang_track_expense/domain/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	List() (users []entity.User, err error)

	GetByUserId(id uuid.UUID) (user entity.User, err error)

	GetByUserName(username string) (user entity.User, err error)

	Create(name string) (err error)
}
