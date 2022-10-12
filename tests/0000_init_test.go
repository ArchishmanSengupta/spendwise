package tests

import (
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ArchishmanSengupta/expense-tracker/api"
	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var Server *httptest.Server

var (
	Id         = ""
	Amount     = ""
	Type       = ""
	created_at = ""
	updated_at = ""
)

func TestMain(m *testing.M) {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file", err.Error())
		return
	}
	// Connect to test db
	cmd.DbConn, err = cmd.Connect()

	dbConn := cmd.DbConn
	if err != nil {
		log.Fatalln("Error connecting to test db: ", err.Error())
		return
	}
	utils.ClearTestDatabase(dbConn)
	// Initialize new router for test
	router := mux.NewRouter()

	api.RegisterRoutes(router)

	// Start new test server
	Server = httptest.NewServer(router)

	// Run tests
	exitVal := m.Run()

	// Close test server
	Server.Close()
	utils.ClearTestDatabase(dbConn)

	// Close test database connection
	_ = dbConn.Close()

	// Exit main test
	os.Exit(exitVal)
}
