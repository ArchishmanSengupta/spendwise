/*******************************************************
 * THIS FILE HOLDS THE LOGIC FOR THE TRANSACTION TABLE *
 *******************************************************/

package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
	"github.com/jmoiron/sqlx"
)

/*
* Transaction Structure of the Transaction model
@struct
@field ID
@field Uuid
@field Amount
@field Type
@field CreatedAt
@field UpdatedAt
*/
type Transaction struct {
	ID        int       `db:"id" json:"id"`
	Uuid      string    `db:"uuid" json:"uuid"`
	Amount    int64     `db:"amount" json:"amount"`
	Type      string    `db:"type" json:"type"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

var TotalTransactionCount int

/*
* This function returns all the transactions in the ElephantSQL database

* @receiver t
* @param typeFromTheUrl
* @param dateFromTheUrl
* @return []*Transaction
* @return error
 */
func (t *Transaction) GetAllTransactions(typeFromTheUrl string, dateFromTheUrl string, limit int, offset int, db *sqlx.DB) ([]*Transaction, error) {
	transactions := make([]*Transaction, 0)

	// execute the select query
	err := cmd.DbConn.Select(&transactions, "SELECT amount,type,created_at, updated_at FROM transaction ORDER BY id LIMIT $1 OFFSET $2", limit, offset)

	totalRecordsCountQuery := `SELECT COUNT(*) FROM transaction`

	err = db.QueryRowx(totalRecordsCountQuery).Scan(&TotalTransactionCount)

	// if an error is found, return that error
	if err != nil {
		fmt.Printf("Error found in GetAllTransactions----->\n", err)
	}
	return transactions, nil
}

/*
*This Retrieve Method is used to retrieve a transaction from the database based on the attributeMap passed
receiver t
param - database
param attributeMap
return *Transaction
return error
*/
func (t *Transaction) Retrieve(db *sqlx.DB, attributeMap map[string]interface{}) (*Transaction, error) {

	query := `SELECT amount, type, created_at, updated_at FROM transaction WHERE `

	// Check for id or uuid in the attributeMap and construct the WHERE clause
	whereClause := ""
	if id, ok := attributeMap["id"]; ok {
		whereClause = fmt.Sprintf("id='%d'", id)
	} else if uuid, ok := attributeMap["uuid"]; ok {
		whereClause = fmt.Sprintf("uuid='%s'", uuid)
	}

	// Append the WHERE clause to the query
	query += whereClause
	// Execute Get operation on the scan table
	err := db.Get(t, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrResourceNotFound
		}
		return nil, err
	}

	return t, nil
}

/**
* * This Insert Method creates a record in the transaction table
* @receiver t
* ? do i need to check the validity of the transaction before inserting it into the database
* ? as it may corrupt the DB
* @param none
* @return created *Transaction, an error if any
 */
func (t *Transaction) Insert() (*Transaction, error) {
GenerateNewUUID:
	uuid := utils.CreateNewUUID()
	transactionInstance := Transaction{}
	dbConn := cmd.DbConn
	_, err := transactionInstance.Retrieve(dbConn, map[string]interface{}{"uuid": uuid})

	if err == nil {
		goto GenerateNewUUID
	}

	t.Uuid = uuid
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	insertQuery := `INSERT INTO transaction (uuid, type, amount, created_at, updated_at) VALUES (:uuid, :type, :amount, :created_at, :updated_at)`

	// execute the insert query
	_, err = cmd.DbConn.NamedExec(insertQuery, &t)

	// if an error is found, return it
	if err != nil {
		fmt.Println("Error found while Inserting---->", err)
		return nil, err
	}

	return t, nil
}

/*
* This UpdateTransaction filter is used to update a transaction

  - @param t
  - @param uuid
  - @return *Transaction
  - @return error
*/
func (t *Transaction) UpdateTransaction(db *sqlx.DB, attributeMap map[string]interface{}) (*Transaction, error) {

	t.UpdatedAt = time.Now()

	query := `UPDATE transaction SET amount=:amount,type=:type `
	whereClause := ""
	if id, ok := attributeMap["id"]; ok {
		whereClause = fmt.Sprintf("WHERE id='%d'", id)
	} else if uuid, ok := attributeMap["uuid"]; ok {
		whereClause = fmt.Sprintf("WHERE uuid='%s'", uuid)
	}

	// Append the WHERE clause to the query
	query += whereClause
	// update a query in the database
	_, err := cmd.DbConn.NamedExec(query, t)

	// if an error is found
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrResourceNotFound
		default:
			return nil, err
		}
	}

	return t, nil
}

// This Delete Method is to delete a transaction from the database based on the uuid.

// @receiver t
// @param db
// @param attributeMap
// @return error
func (t *Transaction) Delete(db *sqlx.DB, attributeMap map[string]interface{}) error {

	query := "DELETE FROM transaction WHERE "
	// Check for id or uuid or status in the attributeMap and construct the WHERE clause
	whereClause := ""

	if id, ok := attributeMap["id"]; ok {
		whereClause = fmt.Sprintf("id=%d", id)
	} else if uuid, ok := attributeMap["uuid"]; ok {
		whereClause = fmt.Sprintf("uuid='%s'", uuid)
	} else if status, ok := attributeMap["status"]; ok {
		whereClause = fmt.Sprintf("status='%s'", status)
	}

	//Append the Where clause to the query
	deleteQuery := query + whereClause

	//Execute the DELETE query
	result, err := db.Exec(deleteQuery)

	//If an error is found, return it
	if err != nil {
		return err
	}

	//Check if any rows were affected
	rowsAffected, err := result.RowsAffected()

	//If an error is found, return it
	if err != nil {
		return err
	}

	//If no rows were affected, return an error
	if rowsAffected == 0 {
		return utils.ErrResourceNotFound
	}
	return nil
}

//This Filter Method is to filter out the query based on the type , amount and date

// @receiver t
// @param db
// @param attributeMap
// @return []*Transaction
// @return error
func (t *Transaction) Filter(db *sqlx.DB, attributeMap map[string]interface{}, limit, offset int) ([]*Transaction, error) {

	whereClause, err := utils.GenerateQueryWhereClause(attributeMap)
	// if any error found in generating the where clause
	if err != nil {
		return nil, err

	}
	query := "SELECT amount, type, created_at FROM transaction "

	filteredQuery := query + whereClause + "LIMIT $1 OFFSET $2"

	transactions := make([]*Transaction, 0)

	err = db.Select(&transactions, filteredQuery, limit, offset)

	// if an error is found, return that erorr
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrResourceNotFound
		default:
			return nil, err
		}
	}
	totalRecordsCountQuery := `SELECT COUNT(*) FROM transaction`

	err = db.QueryRowx(totalRecordsCountQuery).Scan(&TotalTransactionCount)
	if err != nil {
		return nil, err
	}
	return transactions, err
}
