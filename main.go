package main

import (
	"fmt"
	"go-casbin/db"
	"go-casbin/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("::: :: Hello go-casbin")
	env := godotenv.Load()

	if env != nil {
		log.Fatalln("::: :: Error loading .env", env)
	}

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))

	db, _ := db.Connect()

	// Start server
	log.Printf("::: :: Listening on port %v", port)
	log.Fatalln(http.ListenAndServe("localhost"+port, routes.Handler(db)))
}