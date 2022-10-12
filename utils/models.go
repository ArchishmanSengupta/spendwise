package utils

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// This CreateNewUUID Function is used to create a new UUID
// CreateNewUUID
//
//	@return string
func CreateNewUUID() string {
	return uuid.New().String()
}

// GenerateQueryWhereClause This GenerateQueryWhereClause Function is used to generate the query where clause based on the attributeMap passed
//
//	@param attributeMap
//	@return string
//	@return error
func GenerateQueryWhereClause(attributeMap map[string]interface{}) (string, error) {
	if len(attributeMap) == 0 {
		return "", errors.New("empty attribute map")
	}
	creditOrDebit := attributeMap["type"]
	date := attributeMap["date"]
	amount := attributeMap["amount"]

	paramCount := 0
	condition := ""

	if creditOrDebit != nil {
		condition = condition + fmt.Sprintf(`type = '%s' `, creditOrDebit)
		paramCount++
	}

	if amount != nil {
		if paramCount > 0 {
			condition = condition + "AND "
		}
		condition = condition + fmt.Sprintf(`amount = %s `, amount)
		paramCount++
	}

	if date != nil {
		and := ""
		if paramCount > 0 {
			and = "AND "
		}
		condition = condition + and + fmt.Sprintf(`DATE(created_at)='%s'`, date)
	}
	whereClause := fmt.Sprintf(`WHERE %s`, condition)
	return whereClause, nil
}
