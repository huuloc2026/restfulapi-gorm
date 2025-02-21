package main

import (
	"database/sql"
	"log"
	"github.com/huuloc2026/restfulapi-gorm.git/cmd/api"
	"github.com/huuloc2026/restfulapi-gorm.git/db"
)
func main(){
	db,err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	server:= api.NewAPIServer(":8080",nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil {
		log.Fatal((err))
	}
	log.Println("DB: Successfully connected")
}