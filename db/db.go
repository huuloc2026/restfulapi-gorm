package db

import (
	"database/sql"

	"log"

	"github.com/huuloc2026/restfulapi-gorm.git/config"
	_ "github.com/lib/pq"
)


func InitDB() (*sql.DB, error) {
    // dbUser := config.Envs.DB_USER
    // dbPassword := config.Envs.DB_PASSWORD
    // dbPort := config.Envs.DB_PORT
    // dbName := config.Envs.DB_NAME
	dbUrl:= config.Envs.DB_URL

	// dsn := fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable",
	// dbUser, dbPassword, dbPort, dbName,)
	
    db, err := sql.Open("postgres", dbUrl)
    if err != nil {
        log.Fatal(err)
    }
    return db, nil
}