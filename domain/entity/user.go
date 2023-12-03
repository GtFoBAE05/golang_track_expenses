package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidName = errors.New("invalid name")
)

type User struct {
	ID   uuid.UUID
	Name string
}

func NewUser(name string) (user User, err error) {
	if name == "" {
		err = ErrInvalidName
		return
	}

	user = User{
		ID:   uuid.New(),
		Name: name,
	}
	return
}
