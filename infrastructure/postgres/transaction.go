package postgres

import (
	"golang_track_expense/domain/entity"
	"golang_track_expense/domain/valueobject"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	Db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{Db: db}
}

type DbTransaction struct {
	ID           uuid.UUID `db:"id"`
	Amount       int       `db:"amount"`
	CreatedAt    time.Time `db:"created_at"`
	UserID       uuid.UUID `db:"user_id"`
	CategoryName string    `db:"category_name"`
	CategoryType string    `db:"category_type"`
}

func (dbTransaction DbTransaction) ToEntity() entity.Transaction {
	return entity.Transaction{
		ID:        dbTransaction.ID,
		Amount:    dbTransaction.Amount,
		CreatedAt: dbTransaction.CreatedAt,
		UserId:    dbTransaction.UserID,
		Category: valueobject.Category{
			Name: dbTransaction.CategoryName,
			Type: dbTransaction.CategoryType,
		},
	}
}

func (t *TransactionRepository) GetAll() ([]entity.Transaction, error) {
	var dbTransactions []DbTransaction

	query := `SELECT id, amount, created_at, user_id, 
	category_name, category_type FROM transactions`

	err := t.Db.Select(&dbTransactions, query)

	var transactions []entity.Transaction
	for _, dbTransaction := range dbTransactions {
		transactions = append(transactions, dbTransaction.ToEntity())
	}

	return transactions, err
}

func (t *TransactionRepository) GetByTransactionId(id uuid.UUID) (entity.Transaction, error) {
	var dbTransactions DbTransaction

	query := `SELECT id, amount, created_at, user_id, 
	category_name, category_type FROM transactions WHERE id = $1`

	err := t.Db.Get(&dbTransactions, query, id)

	var transaction = dbTransactions.ToEntity()

	return transaction, err
}

func (t *TransactionRepository) GetByUserId(id uuid.UUID) ([]entity.Transaction, error) {
	var dbTransactions []DbTransaction

	query := `SELECT id, amount, created_at, user_id, 
	category_name, category_type FROM transactions WHERE user_id = $1`

	err := t.Db.Select(&dbTransactions, query, id)

	var transactions []entity.Transaction
	for _, dbTransaction := range dbTransactions {
		transactions = append(transactions, dbTransaction.ToEntity())
	}

	return transactions, err
}

func (t *TransactionRepository) CreateOrder(transaction entity.Transaction) error {

	query := `INSERT INTO transactions (id,  amount, created_at, user_id, category_name, category_type) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := t.Db.Exec(query, transaction.ID, transaction.Amount, transaction.CreatedAt, transaction.UserId,
		transaction.Category.Name, transaction.Category.Type)

	return err
}
