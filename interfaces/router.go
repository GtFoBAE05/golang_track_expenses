package interfaces

import "github.com/gin-gonic/gin"

func SetRoutes(router *gin.Engine, userHandlers *UserHandlers, tranTransactionHandler *TransactionHandler) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("/", userHandlers.ListUsers)
		userRouter.GET("/:id", userHandlers.GetUserByID)
		userRouter.POST("/", userHandlers.CreateUser)
	}

	transactionRouter := router.Group("/transaction")
	{
		transactionRouter.GET("/", tranTransactionHandler.ListTransactions)
		transactionRouter.GET("/:id", tranTransactionHandler.GetTransactionByID)
		transactionRouter.GET("/user/:id", tranTransactionHandler.GetTransactionByUserID)
		transactionRouter.POST("/", tranTransactionHandler.CreateTransaction)
	}
}
