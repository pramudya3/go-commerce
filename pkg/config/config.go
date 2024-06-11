package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	TimeoutCtx  int
	ServerAddr  string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DBName      string
	RedisHost   string
	RedisPass   string
	TokenSecret string
}

func LoadConfig() *Config {
	godotenv.Load()

	// default setting
	cfg := &Config{
		TimeoutCtx: 5,
		ServerAddr: ":8080",
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPass:     "postgres",
		DBName:     "postgres",
		RedisHost:  "localhost:6379",
	}

	if timeoutCtx, ok := os.LookupEnv("TIMEOUT_CTX"); ok && timeoutCtx != "" {
		timeout, err := strconv.Atoi(timeoutCtx)
		if err != nil {
			log.Printf("failed parsing timeout context, err: %v\n", err)
			return nil
		}
		cfg.TimeoutCtx = timeout
	}

	if serverAddr, ok := os.LookupEnv("SERVER_ADDR"); ok && serverAddr != "" {
		cfg.ServerAddr = serverAddr
	}

	if dbHost, ok := os.LookupEnv("DB_HOST"); ok && dbHost != "" {
		cfg.DBHost = dbHost
	}

	if dbPort, ok := os.LookupEnv("DB_PORT"); ok && dbPort != "" {
		cfg.DBPort = dbPort
	}

	if dbUser, ok := os.LookupEnv("DB_USER"); ok && dbUser != "" {
		cfg.DBUser = dbUser
	}

	if dbPass, ok := os.LookupEnv("DB_PASS"); ok && dbPass != "" {
		cfg.DBPass = dbPass
	}

	if dbName, ok := os.LookupEnv("DB_NAME"); ok && dbName != "" {
		cfg.DBName = dbName
	}

	if redisHost, ok := os.LookupEnv("REDIS_HOST"); ok && redisHost != "" {
		cfg.RedisHost = redisHost
	}

	if redisPass, ok := os.LookupEnv("REDIS_PASS"); ok && redisPass != "" {
		cfg.RedisPass = redisPass
	}

	if tokenSecret, ok := os.LookupEnv("TOKEN_SECRET"); ok && tokenSecret != "" {
		cfg.TokenSecret = tokenSecret
	}

	return cfg
}
