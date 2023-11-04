package todo

import (
	"context"
	"log"
	"os"
	"testing"
	"todo/biz/testcontainers"
)

func TestMain(m *testing.M) {
	container, err := testcontainers.CreatePostgresContainer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := container.Terminate(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()

	code := m.Run()
	os.Exit(code)
}
