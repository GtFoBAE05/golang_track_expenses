package interfaces

import (
	"golang_track_expense/application"
	"golang_track_expense/domain/entity"
	"golang_track_expense/domain/valueobject"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionService *application.TransactionService
}

func NewTransactionHandler(transactionService *application.TransactionService) *TransactionHandler {
	return &TransactionHandler{TransactionService: transactionService}
}

func (th *TransactionHandler) ListTransactions(c *gin.Context) {
	transaction, err := th.TransactionService.List()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, transaction)
}

func (th *TransactionHandler) GetTransactionByID(c *gin.Context) {
	transaction, err := th.TransactionService.GetByTransactionId(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, transaction)
}

func (th *TransactionHandler) GetTransactionByUserID(c *gin.Context) {
	transaction, err := th.TransactionService.GetByUserId(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, transaction)
}

func (th *TransactionHandler) CreateTransaction(c *gin.Context) {
	var body TransactionRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	category, err := valueobject.NewCategory(body.CategoryName, body.CategoryType)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	transaction, err := entity.NewTransaction(body.Amount, body.UserId, category)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = th.TransactionService.Create(transaction)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Transaction created successfully!"})
}

type TransactionRequest struct {
	UserId       string `json:"user_id"`
	Amount       int    `json:"amount"`
	CategoryName string `json:"category_name"`
	CategoryType string `json:"category_type"`
}
