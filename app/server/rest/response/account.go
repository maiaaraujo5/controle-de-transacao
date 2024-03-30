package response

import "github.com/maiaaraujo5/controle-de-transacao/app/domain/model"

type Account struct {
	ID             int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func NewAccount(domain *model.Account) *Account {
	return &Account{
		ID:             domain.ID,
		DocumentNumber: domain.DocumentNumber,
	}
}
