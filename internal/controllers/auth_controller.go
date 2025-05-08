package controllers

import (
	"net/http"

	"github.com/gabrielagui373/obiwanapp-api/internal/models"
	"github.com/gabrielagui373/obiwanapp-api/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authController *services.AuthService) *AuthController {
	return &AuthController{authService: authController}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	var existingUser models.User
	if err := c.authService.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email already in use"})
		return
	}

	// Hash password
	hashedPassword, err := c.authService.HashPassword(input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
		return
	}

	//Create User
	user := models.User{
		Email:        input.Email,
		PasswordHash: hashedPassword,
		IsActive:     true,
	}

	if err := c.authService.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := c.authService.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := c.authService.RefreshToken(input.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (u *AuthController) ProtectedRoute(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": "access granted", "user": user})
}
