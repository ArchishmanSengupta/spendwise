package models

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ArchishmanSengupta/expense-tracker/cmd"
)

/*
Structure of the Transaction model
*/
type Transaction struct {
	ID        int       `db:"id" json:"id"`
	Uuid      string    `db:"uuid" json:"uuid"`
	Amount    *big.Int  `db:"amount" json:"amount"`
	Type      string    `db:"type" json:"type"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func GetAllTransactions(transaction *Transaction) ([]Transaction, error) {
	transactions := make([]Transaction, 0)

	// execute the select query
	err := cmd.DbConn.Select(&transactions, "SELECT * FROM transactions")

	// if an error is found, return that error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func CreateTransaction(transaction *Transaction, id int) ([]Transaction, error) {
	transactions := make([]Transaction, 0)
	insertQuery := `INSERT INTO transactions (uuid, id, type, amount, created_at, updated_at) VALUES (:uuid, :id, :type, :amount, :created_at, :updated_at)`

	// execute the insert query
	_, err := cmd.DbConn.NamedExec(insertQuery, &transaction)

	// if an error is found, return it
	if err != nil {
		fmt.Println("Error found", err)
	}

	return transactions, nil
}
func Retrieve(transaction *Transaction, id int) (*Transaction, error) {
	// execute the query
	err := cmd.DbConn.Get(transaction, "SELECT * FROM transactions WHERE id = $1 LIMIT 1", id)

	// if an error is found, return it
	if err != nil {
		fmt.Println("Error found", err)
	}

	return transaction, err
}
func UpdateTransaction(transaction *Transaction, id int) ([]Transaction, error) {
	transactions := make([]Transaction, 0)

	updateStmt := `UPDATE transactions SET amount=:amount, created_at=:created_at, updated_at=:updated_at,type=:type WHERE id=:id`

	// update a query in the database
	_, err := cmd.DbConn.NamedExec(updateStmt, transaction)

	// if an error is found
	if err != nil {
		fmt.Println("Error found", err)
	}

	return transactions, nil
}

func DeleteTransaction(transaction *Transaction, id int) error {

	deleteStmt := `DELETE FROM transaction WHERE id=$1`

	// check if the record with this id exists in the transactions table
	err := cmd.DbConn.Get(&transaction, "SELECT * FROM transactions WHERE id=$1 LIMIT 1", id)

	// if an error is found, return it
	if err != nil {
		fmt.Println("Error found", err)
	}
	// execute the delete query
	_, err = cmd.DbConn.Exec(deleteStmt, id)

	// if an error is found, return it
	if err != nil {
		fmt.Println("Error found", err)
	}

	return nil
}
