package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	defaultHost     = "localhost"
	defaultPort     = 5432
	defaultUser     = "postgres"
	defaultPassword = "password"
	defaultDBName   = "postgres"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func getDBConfig() DBConfig {
	return DBConfig{
		Host:     getEnv("DB_HOST", defaultHost),
		Port:     getEnvAsInt("DB_PORT", defaultPort),
		User:     getEnv("DB_USER", defaultUser),
		Password: getEnv("DB_PASSWORD", defaultPassword),
		DBName:   getEnv("DB_NAME", defaultDBName),
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func SetupDB() *sql.DB {
	dbConfig := getDBConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func MigrateDB() {
	dbConfig := getDBConfig()
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.User,
		dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	m, err := migrate.New("file://migrations", connectionString)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		panic(err)
	}
}
