package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"todo/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(dbConfig config.DbConfig, path string) error {
	if err := resolveTemplates(path); err != nil {
		return err
	}

	connectionString := dbConfig.GetConnectionString()

	basePath := fmt.Sprintf("file://%s/", path+"/resolved")
	m, err := migrate.New(basePath, connectionString)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func SetupDb(config config.DbConfig) (*sql.DB, error) {
	connectionString := config.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func resolveTemplates(path string) error {
	files, err := getFiles(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := resolveFile(file, path); err != nil {
			return err
		}
	}

	return nil
}

func getFiles(path string) ([]string, error) {
	dir, err := os.ReadDir(path + "/templates")
	if err != nil {
		return nil, err
	}
	var files []string
	for _, file := range dir {
		println(file.Name())
		files = append(files, file.Name())
	}
	return files, nil
}

func resolveFile(file string, path string) error {
	// read file
	contentBytes, err := os.ReadFile(filepath.Join(path, "templates", file))
	if err != nil {
		return err
	}

	// create new template
	tmpl, err := template.New(file).Parse(string(contentBytes))
	if err != nil {
		return err
	}

	// delete if resolved directory exists
	if _, err := os.Stat(filepath.Join(path, "resolved")); err == nil {
		os.RemoveAll(filepath.Join(path, "resolved"))
	}

	// create resolved directory if it doesn't exist
	if _, err := os.Stat(filepath.Join(path, "resolved")); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(path, "resolved"), 0755)
	}

	// create new file to write resolved template
	resolvedFile, err := os.Create(filepath.Join(path, "resolved", file))
	if err != nil {
		return err
	}
	defer resolvedFile.Close()

	// create a map of data to replace in the template
	data := map[string]string{
		"User":     "appuser",
		"Password": "password",
	}

	// execute template with data and write to file
	err = tmpl.Execute(resolvedFile, data)
	if err != nil {
		return err
	}
	return nil
}
