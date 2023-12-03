package postgres

import (
	"golang_track_expense/domain/entity"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	Db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{Db: db}
}

func (t *TransactionRepository) GetAll() ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	query := `SELECT id, name FROM transactions`

	err := t.Db.Select(&transactions, query)

	return transactions, err
}

func (t *TransactionRepository) GetByTransactionId(id uuid.UUID) (entity.Transaction, error) {
	var transaction entity.Transaction

	query := `SELECT id, name FROM transactions WHERE id = $1`

	err := t.Db.Get(&transaction, query, id)

	return transaction, err
}

func (t *TransactionRepository) GetByUserId(id uuid.UUID) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	query := `SELECT id, name FROM transactions WHERE user_id = $1`

	err := t.Db.Select(&transactions, query, id)

	return transactions, err
}

func (t *TransactionRepository) CreateOrder(transaction entity.Transaction) error {
	query := `INSERT INTO transactions (id,  amount, created_at, user_id, categori_name, category_type) 
	VALUES ($1, $2, $3, $4, $5, $6`

	_, err := t.Db.Exec(query, transaction.ID, transaction.Amount, transaction.CreatedAt, transaction.UserId,
		transaction.Category.Name, transaction.Category.Type)

	return err
}
