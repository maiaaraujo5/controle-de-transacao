package errors

import "errors"

type accountNotExists struct {
	Message string
}

func AccountNotExists(message string) error {
	return &accountNotExists{
		message,
	}
}

func (a *accountNotExists) Error() string {
	return a.Message
}

func IsAccountNotExists(err error) bool {
	var myErr *accountNotExists
	return errors.As(err, &myErr)
}
