package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/huuloc2026/restfulapi-gorm.git/cmd/api"
	"github.com/huuloc2026/restfulapi-gorm.git/db"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	initStorage(database)

	server := api.NewAPIServer(":8080", database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("OKe")
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal((err))
	}
	log.Println("DB: Successfully connected")
}
