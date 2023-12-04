package repository

import (
	"golang_track_expense/domain/entity"

	"github.com/google/uuid"
)

type TransactionRepository interface {
	GetAll() (transactions []entity.Transaction, err error)

	GetByTransactionId(transactionId uuid.UUID) (transaction entity.Transaction, err error)

	GetByUserId(userId uuid.UUID) (transactions []entity.Transaction, err error)

	CreateOrder(entity.Transaction) (err error)
}
