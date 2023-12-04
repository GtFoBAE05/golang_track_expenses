package application

import (
	"golang_track_expense/domain/entity"
	"golang_track_expense/domain/repository"

	"github.com/google/uuid"
)

type TransactionService struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) *TransactionService {
	return &TransactionService{TransactionRepository: transactionRepository}
}

func (t *TransactionService) List() ([]entity.Transaction, error) {
	return t.TransactionRepository.GetAll()
}

func (t *TransactionService) GetByTransactionId(id string) (entity.Transaction, error) {
	var transaction entity.Transaction
	uid, err := uuid.Parse(id)
	if err != nil {
		return transaction, err
	}
	return t.TransactionRepository.GetByTransactionId(uid)
}

func (t *TransactionService) GetByUserId(id string) ([]entity.Transaction, error) {
	var transaction []entity.Transaction
	uid, err := uuid.Parse(id)
	if err != nil {
		return transaction, err
	}
	return t.TransactionRepository.GetByUserId(uid)
}

func (t *TransactionService) Create(transaction entity.Transaction) error {

	return t.TransactionRepository.CreateOrder(transaction)
}
