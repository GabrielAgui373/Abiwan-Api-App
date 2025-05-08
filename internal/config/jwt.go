package config

import "time"

type JWTConfig struct {
	SecretKey       string
	AccessTokenExp  time.Duration
	RefreshTokenExp time.Duration
}

func LoadJWTConfig() *JWTConfig {
	return &JWTConfig{
		SecretKey:       getEnv("JWT_SECRET", "very-secret-key"),
		AccessTokenExp:  time.Hour * 1,
		RefreshTokenExp: time.Hour * 24 * 7,
	}
}


