package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gabrielagui373/obiwanapp-api/internal/config"
	"github.com/gabrielagui373/obiwanapp-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB        *gorm.DB
	JWTConfig *config.JWTConfig
}

func NewAuthService(DB *gorm.DB, JWTConfig *config.JWTConfig) *AuthService {
	return &AuthService{DB: DB, JWTConfig: JWTConfig}
}

// hashes a password using bcrypt
func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// compares a password with its hash
func (s *AuthService) CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *AuthService) GenerateJWT(user *models.User, tokenType string) (string, error) {
	var expirationTime time.Time

	if tokenType == "access" {
		expirationTime = time.Now().Add(s.JWTConfig.AccessTokenExp)
	} else {
		expirationTime = time.Now().Add(s.JWTConfig.RefreshTokenExp)
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     expirationTime.Unix(),
		"type":    tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.JWTConfig.SecretKey))
}

// validates a JWT token - verify
func (s *AuthService) ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.JWTConfig.SecretKey), nil
	})
}

// authenticates a user and return tokens
func (s *AuthService) Login(email, password string) (map[string]string, error) {
	var user models.User
	if err := s.DB.Preload("Roles.Permission").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("invalid credencials")
	}

	if !s.CheckPassword(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return nil, errors.New("account is not active")
	}

	accessToken, err := s.GenerateJWT(&user, "access")
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.GenerateJWT(&user, "refresh")
	if err != nil {
		return nil, err
	}

	refreshTokenRecord := models.Token{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(s.JWTConfig.RefreshTokenExp),
		TokenType: "refresh",
	}

	if err := s.DB.Create(&refreshTokenRecord).Error; err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

// generates new access token using refresh token
func (s *AuthService) RefreshToken(refreshToken string) (map[string]string, error) {
	token, err := s.ValidateJWT(refreshToken)

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	if claims["type"] != "refresh" {
		return nil, errors.New("invalid token type")
	}

	var tokenRecord models.Token
	if err := s.DB.Where("token = ? AND is_revoked = ?", refreshToken, false).First(&tokenRecord).Error; err != nil {
		return nil, errors.New("token not found or revoked")
	}

	if time.Now().After(tokenRecord.ExpiresAt) {
		return nil, errors.New("token expired")
	}

	var user models.User
	if err := s.DB.Preload("Roles.Permissions").First(&user, claims["user_id"]).Error; err != nil {
		return nil, errors.New("user not found")
	}

	newAccessToken, err := s.GenerateJWT(&user, "access")
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token": newAccessToken,
	}, nil
}
