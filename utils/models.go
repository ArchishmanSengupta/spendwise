package utils

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func CreateNewUUID() string {
	return uuid.New().String()
}

func GenerateQueryWhereClause(attributeMap map[string]interface{}) (string, error) {
	if len(attributeMap) == 0 {
		return "", errors.New("empty attribute map")
	}

	condition := ""
	for key, value := range attributeMap {
		if len(attributeMap) == 1 {
			// Check if the value is string int or slice and construct the query accordingly.
			condition = condition + fmt.Sprintf(`%s = '%v'`, key, value)
			delete(attributeMap, key)

		} else {
			// If there are more than one attribute in the map, then we need to add AND
			condition = condition + fmt.Sprintf(`%s = '%v' AND `, key, value)
			delete(attributeMap, key)
		}
	}
	whereClause := fmt.Sprintf(`WHERE %s`, condition)
	return whereClause, nil
}
