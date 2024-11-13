package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/pp00x/hydrate/internal/model"
    "github.com/pp00x/hydrate/internal/repository"
    "github.com/pp00x/hydrate/internal/service"
    "github.com/pp00x/hydrate/pkg/util"
    "gorm.io/gorm"
)

var userService service.UserService

func init() {
    db := util.GetDB()
    userRepo := repository.NewUserRepository(db)
    userService = service.NewUserService(userRepo)
}

type RegisterInput struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}

func Register(c *gin.Context) {
    var input RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := model.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: input.Password,
    }

    if err := userService.Register(&user); err != nil {
        if err == gorm.ErrDuplicatedKey {
            c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

type LoginInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
    var input LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := userService.Login(input.Email, input.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := service.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}