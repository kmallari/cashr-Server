package db

import (
	"context"
	"database/sql"
	db "github.com/kmallari/cashr-Server/db/postgres"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func ConnectDb() *sql.DB {
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	// Configure pool settings (optional)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	return db
}

func AttachQueries(databaseConnection *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(databaseConnection)
		ctx := context.WithValue(r.Context(), "queries", queries)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AttachDB(database *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "database", database)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}