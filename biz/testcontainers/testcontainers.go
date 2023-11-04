package testcontainers

import (
	"context"
	"time"
	"todo/config"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func CreatePostgresContainer() (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "test",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(time.Duration(5 * time.Second)),
	}
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	return container, nil
}

func GetMigrationTestConfig(container testcontainers.Container) (config.DbConfig, error) {
	host, err := container.Host(context.Background())
	if err != nil {
		return config.DbConfig{}, err
	}
	port, err := container.MappedPort(context.Background(), "5432")
	if err != nil {
		return config.DbConfig{}, err
	}
	return config.DbConfig{
		Host:     host,
		Port:     port.Int(),
		User:     "test",
		Password: "test",
		DbName:   "test",
	}, nil
}

func GetAppTestConfig(migrationConfig config.DbConfig) config.DbConfig {
	return config.DbConfig{
		Host:     migrationConfig.Host,
		Port:     migrationConfig.Port,
		User:     "appuser",
		Password: "password",
		DbName:   "test",
	}
}
