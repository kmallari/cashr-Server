package db

import (
	"database/sql"
	"github.com/joho/godotenv"
	sqlc "github.com/kmallari/cashr-Server/db/postgres"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var Client *sql.DB
var Queries *sqlc.Queries

func init() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	Client = ConnectDb()
	Queries = AttachQueries(Client)
}

func ConnectDb() *sql.DB {
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	// Configure pool settings (optional)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	log.Println("Initializing database...", db.Ping())
	return db
}

func AttachQueries(databaseConnection *sql.DB) *sqlc.Queries {
	log.Println("Initializing queries...")
	return sqlc.New(databaseConnection)
}
