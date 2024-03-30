package errors

import "errors"

type operationNotValid struct {
	Message string
}

func OperationNotValid(message string) error {
	return &operationNotValid{
		message,
	}
}

func (a *operationNotValid) Error() string {
	return a.Message
}

func IsOperationNotValid(err error) bool {
	var myErr *operationNotValid
	return errors.As(err, &myErr)
}
