package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/huuloc2026/restfulapi-gorm.git/db"
)

func main() {
	// Initialize database connection
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Database initialization error:", err)
	}

	instance, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		log.Fatal("Postgres instance error:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", // Fixed the typo
		"postgres",
		instance,
	)
	if err != nil {
		log.Fatal("Migration initialization error:", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration command: 'up' or 'down'")
	}

	cmd := os.Args[1]
	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration up error:", err)
		} else {
			log.Println("Migration applied successfully.")
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration down error:", err)
		} else {
			log.Println("Migration rolled back successfully.")
		}
	default:
		log.Fatal("Unknown command. Use 'up' or 'down'.")
	}
}
