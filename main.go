package main

import (
	"fmt"
	db "go-technical-test-synopsis/database"
	"go-technical-test-synopsis/src/web/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	fmt.Println(os.Getenv("MYSQL_CONNECTION"))
	dbMysql := db.InitMysql(os.Getenv("MYSQL_CONNECTION"))

	server.Run(dbMysql)
}
