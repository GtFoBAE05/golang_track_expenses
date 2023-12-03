package entity

import (
	"errors"
	"golang_track_expense/domain/valueobject"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyAmount = errors.New("amount cannot be empty")
)

type Transaction struct {
	ID        uuid.UUID
	Amount    int
	CreatedAt time.Time

	UserId uuid.UUID

	Category valueobject.Category
}

func NewTransaction(amount int, userId string, category valueobject.Category) (trx Transaction, err error) {
	if amount == 0 {
		err = ErrEmptyAmount
		return
	}

	userUid, err := uuid.Parse(userId)
	if err != nil {
		return
	}

	trx = Transaction{
		ID:        uuid.New(),
		Amount:    amount,
		UserId:    userUid,
		CreatedAt: time.Now(),
		Category:  category,
	}
	return
}
