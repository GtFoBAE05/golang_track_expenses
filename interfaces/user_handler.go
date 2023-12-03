package interfaces

import (
	"errors"
	"golang_track_expense/application"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandlers struct {
	UserService *application.UserService
}

func NewUserHandler(userService *application.UserService) *UserHandlers {
	return &UserHandlers{UserService: userService}
}

func (uh *UserHandlers) ListUsers(c *gin.Context) {
	users, err := uh.UserService.List()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

func (uh *UserHandlers) GetUserByID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uh.UserService.GetByUserId(userID.String())
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (uh *UserHandlers) CreateUser(c *gin.Context) {
	var body application.UserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	err := uh.UserService.Create(body.Name)
	if err != nil {
		if errors.Is(err, errors.New("user already exists")) {
			c.JSON(409, gin.H{"error": err.Error()})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully!"})
}
