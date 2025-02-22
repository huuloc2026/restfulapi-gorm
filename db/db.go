package db

import (
	"database/sql"
	"fmt"

	"log"

	"github.com/huuloc2026/restfulapi-gorm.git/config"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	// dbUser := config.Envs.DB_USER
	// dbPassword := config.Envs.DB_PASSWORD
	// dbPort := config.Envs.DB_PORT
	// dbName := config.Envs.DB_NAME
	dbUrl := config.Envs.DB_URL

	// dsn := fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable",
	// dbUser, dbPassword, dbPort, dbName,)

	db, err := sql.Open("postgres", dbUrl)
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        role VARCHAR(50) DEFAULT 'user',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

	CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description VARCHAR(255) NOT NULL UNIQUE,
        categories VARCHAR(255) DEFAULT 'electronic',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
	
	
	`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	fmt.Println("Database initialized successfully.")
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
