package repository

import (
	"golang_track_expense/domain/aggregate"
	"golang_track_expense/domain/entity"

	"github.com/google/uuid"
)

type TransactionRepository interface {
	GetAll() (histories []aggregate.History, err error)

	GetByTransactionId(transactionId uuid.UUID) (history aggregate.History, err error)

	GetByUserId(userId uuid.UUID) (history []aggregate.History, err error)

	CreateOrder(trx entity.Transaction, user entity.User) (err error)
}
