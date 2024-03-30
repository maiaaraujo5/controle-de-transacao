package errors

import "errors"

type AccountAlreadyExistss struct {
	Message string
}

func AccountAlreadyExists(message string) error {
	return &AccountAlreadyExistss{
		message,
	}
}

func (a AccountAlreadyExistss) Error() string {
	return a.Message
}

func IsAccountAlreadyExists(err error) bool {
	return errors.Is(err, new(AccountAlreadyExistss))
}
