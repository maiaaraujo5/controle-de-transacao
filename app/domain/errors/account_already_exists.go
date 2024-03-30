package errors

import "errors"

type accountAlreadyExists struct {
	Message string
}

func AccountAlreadyExists(message string) error {
	return &accountAlreadyExists{
		message,
	}
}

func (a *accountAlreadyExists) Error() string {
	return a.Message
}

func IsAccountAlreadyExists(err error) bool {
	var myErr *accountAlreadyExists
	return errors.As(err, &myErr)
}
