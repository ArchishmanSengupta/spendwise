package models

import (
	"fmt"
	"time"

	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
)

/*
Structure of the Transaction model
*/
type Transaction struct {
	ID        int       `db:"ID" json:"id"`
	Uuid      string    `db:"Uuid" json:"uuid"`
	Amount    int64     `db:"Amount" json:"amount"`
	Type      string    `db:"Type" json:"type"`
	CreatedAt time.Time `db:"CreatedAt" json:"createdAt"`
	UpdatedAt time.Time `db:"UpdatedAt" json:"updatedAt"`
}

func GetAllTransactions(transaction *Transaction) ([]Transaction, error) {
	transactions := make([]Transaction, 0)

	// execute the select query
	err := cmd.DbConn.Select(&transactions, "SELECT * FROM transaction")

	// if an error is found, return that error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (transaction *Transaction) GetByUuid(uuid string) (*Transaction, error) {
	// execute the query
	query := `SELECT * FROM transaction WHERE Uuid = $1`

	// testing := fmt.Sprintf(query, uuid)

	// fmt.Println("Testing---->", testing)

	err := cmd.DbConn.Get(transaction, query, uuid)

	// if an error is found, return it
	if err != nil {
		fmt.Println("Error found in GetByUuid Model---->", err)
	}

	return transaction, err
}

func (transaction *Transaction) Insert() (*Transaction, error) {
GenerateNewUUID:
	uuid := utils.CreateNewUUID()
	// fmt.Println(transaction)

	transactionInstance := Transaction{}

	txn, err := transactionInstance.GetByUuid(uuid)

	if err == nil {
		goto GenerateNewUUID
	}

	transaction.Uuid = uuid
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	fmt.Println("Transaction------->", transaction)
	insertQuery := `INSERT INTO transaction (Uuid, Type, Amount, CreatedAt, UpdatedAt) VALUES (:Uuid, :Type, :Amount, :CreatedAt, :UpdatedAt)`

	// execute the insert query
	_, err = cmd.DbConn.NamedExec(insertQuery, &transaction)

	// if an error is found, return it
	if err != nil {
		fmt.Println("Error found while Inserting---->", err)
	}

	// txn, err := transactionInstance.GetByUuid(uuid)

	if err != nil {
		fmt.Println("Error getting transaction by uuid----->", err)
		return nil, err
	}
	return txn, nil
}
func (transaction *Transaction) Retrieve(uuid string) (*Transaction, error) {
	txn := make([]*Transaction, 0)
	// execute the query
	err := cmd.DbConn.Select(&txn, "SELECT * FROM transaction WHERE Uuid = $1 LIMIT 1", uuid)

	// if an error is found, return it
	if err != nil {
		fmt.Println("Error found", err)
	}

	return txn[0], err
}

func Retrieve(transaction *Transaction, uuid string) (*Transaction, error) {
	// execute the query
	err := cmd.DbConn.Get(transaction, "SELECT * FROM transaction WHERE Uuid = $1 LIMIT 1", uuid)

	// if an error is found, return it
	if err != nil {
		fmt.Println("Error found", err)
	}

	return transaction, err
}

func UpdateTransaction(transaction *Transaction, uuid string) (*Transaction, error) {
	// transactions := make([]Transaction, 0)

	updateStmt := `UPDATE transactions SET Amount=:Amount, CreatedAt=:CreatedAt, UpdatedAt=:UpdatedAt,Type=:Type WHERE ID=:ID`

	// update a query in the database
	_, err := cmd.DbConn.NamedExec(updateStmt, transaction)

	// if an error is found
	if err != nil {
		fmt.Println("Error found", err)
	}

	return transaction, nil
}

func DeleteTransaction(transaction *Transaction, uuid string) (*Transaction, error) {

	deleteStmt := `DELETE FROM transaction WHERE Uuid=$1`

	// check if the record with this id exists in the transactions table
	err := cmd.DbConn.Get(&transaction, "SELECT * FROM transaction WHERE Uuid=$1 LIMIT 1", uuid)

	// if an error is found, return it
	if err != nil {
		fmt.Println("Error found", err)
	}
	// execute the delete query
	_, err = cmd.DbConn.Exec(deleteStmt, uuid)

	// if an error is found, return it
	if err != nil {
		fmt.Println("Error found", err)
	}

	return transaction, nil
}
