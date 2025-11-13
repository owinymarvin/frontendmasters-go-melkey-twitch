package store

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"io/fs"
	"github.com/pressly/goose/v3"

)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}
	fmt.Println("Connected to the Database .... ")
	return db, nil
}

func MigrateFS(db *sql.DB, migrationFS fs.FS, dir string) error {
	goose.SetBaseFS(migrationFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()
	return Migrate(db, dir)
}


func Migrate (db *sql.DB, dir string) error {
	err:= goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate: set dialect %w", err)
	}
	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("migrate: goose up %w", err)
	}


	return nil
}