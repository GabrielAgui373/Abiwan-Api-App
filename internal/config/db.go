package config

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBName:     getEnv("DB_NAME", "studydb"),
		DBPassword: getEnv("DB_PASSWORD", "qwe123"),
	}
}
