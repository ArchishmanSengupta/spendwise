/*********************************************************************
 *       THIS FILE HOLDS THE SERIALIZERS OF THE TRANSACTIONS.        *
 * IT IS USED TO SERIALIZE THE DATA BEFORE SENDING IT TO THE CLIENT. *
 *********************************************************************/

package serializers

import "github.com/ArchishmanSengupta/expense-tracker/api/models"

type TransactionSerializer struct {
	Transactions []*models.Transaction
	Many         bool
}

func (serializer *TransactionSerializer) Serialize() map[string]interface{} {
	serializedData := make(map[string]interface{})

	transactionsArray := make([]interface{}, 0)
	for _, transaction := range serializer.Transactions {
		transactionsArray = append(transactionsArray, map[string]interface{}{
			"type":       transaction.Type,
			"amount":     transaction.Amount,
			"created_at": transaction.CreatedAt,
			"updated_at": transaction.UpdatedAt,
		})
	}

	if serializer.Many {
		serializedData["data"] = transactionsArray
	} else {
		if len(transactionsArray) != 0 {
			serializedData["data"] = transactionsArray[0]
		} else {
			serializedData["data"] = make(map[string]interface{})
		}
	}

	return serializedData
}
