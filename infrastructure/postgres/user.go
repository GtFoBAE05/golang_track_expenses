package postgres

import (
	"errors"
	"golang_track_expense/domain/entity"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	ErrUserAlreadyExist = errors.New("user already exists")
	ErrUserNotFound     = errors.New("user not found")
)

type UserRepository struct {
	Db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (u *UserRepository) GetByUserId(id uuid.UUID) (entity.User, error) {
	var user entity.User

	query := `SELECT id, name FROM users WHERE id = $1`

	err := u.Db.Get(&user, query, id)

	if err != nil {
		return user, ErrUserNotFound
	}

	return user, nil
}

func (u *UserRepository) List() ([]entity.User, error) {
	var users []entity.User

	query := `SELECT id, name FROM users`

	err := u.Db.Select(&users, query)

	return users, err
}

func (u *UserRepository) Create(name string) error {

	_, err := u.GetByUserName(name)
	if err == nil {
		return err
	}

	user, err := entity.NewUser(name)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (id, name) VALUES ($1, $2)`

	_, err = u.Db.Exec(query, user.ID, user.Name)

	return err
}

func (u *UserRepository) GetByUserName(username string) (entity.User, error) {
	var user entity.User

	query := `SELECT id, name FROM users WHERE name = $1`

	err := u.Db.Get(&user, query, username)

	return user, err
}
