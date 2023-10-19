package app

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"

	_ "github.com/lib/pq"
)

type App struct {
	db *sqlx.DB
}

func NewApp() *App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := getDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return &App{
		db: db,
	}
}

func getDB(uri string) (*sqlx.DB, error) {
	// before : directly using sqlx
	// DB, err = sqlx.Connect("postgres", uri)
	// after : using pgx to setup connection
	DB, err := createPGX(uri)
	if err != nil {
		return nil, err
	}
	DB.SetMaxIdleConns(2)
	DB.SetMaxOpenConns(4)
	DB.SetConnMaxLifetime(time.Duration(30) * time.Minute)

	return DB, nil
}

func createPGX(uri string) (*sqlx.DB, error) {
	connConfig, _ := pgx.ParseConfig(uri)
	afterConnect := stdlib.OptionAfterConnect(func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(ctx, `SELECT 1`)
		if err != nil {
			return err
		}
		return nil
	})

	pgxdb := stdlib.OpenDB(*connConfig, afterConnect)
	return sqlx.NewDb(pgxdb, "pgx"), nil
}
