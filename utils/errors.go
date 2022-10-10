package utils

import (
	"strconv"
)

type CustomError struct {
	Message string
	Code    int
}

func (c CustomError) Error() string {
	return c.Message + " " + strconv.Itoa(c.Code)
}

func (c CustomError) AmountError() string {
	return c.Message + " " + strconv.Itoa(c.Code)
}

var ErrResourceNotFound = CustomError{"the specified resource was not found", -1}
var ErrInvalidAmount = CustomError{"the specified amount is invalid", -2}
