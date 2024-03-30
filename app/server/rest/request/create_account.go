package request

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
)

type CreateAccount struct {
	DocumentNumber string `json:"document_number"`
}

func NewCreateAccount(e echo.Context) (*CreateAccount, error) {
	c := new(CreateAccount)
	err := e.Bind(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *CreateAccount) ToDomainModel() *model.Account {
	return &model.Account{
		DocumentNumber: c.DocumentNumber,
	}
}
