/******************************************************
 *       THIS FILE IS THE MAIN EXECUTABLE FILE.       *
 * IT CONNECTS TO THE DATABASE AND STARTS THE SERVER. *
 *    THE REGISTER ROUTES THEN ASSIGNS THE ROUTES.     *
 ******************************************************/

package main

import (
	"log"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api"
	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	r := mux.NewRouter()
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err.Error())
	}

	dbConn, err := cmd.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}
	cmd.DbConn = dbConn
	api.RegisterRoutes(r)

	log.Println("Starting up the server")
	if err := http.ListenAndServe("localhost:3000", r); err != nil {
		log.Panic(err)
	}
}
