package interfaces

import "github.com/gin-gonic/gin"

func SetUserRoutes(router *gin.Engine, userHandlers *UserHandlers) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("/", userHandlers.ListUsers)
		userRouter.GET("/:id", userHandlers.GetUserByID)
		userRouter.POST("/", userHandlers.CreateUser)
	}
}
