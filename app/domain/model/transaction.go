package model

import "time"

type Transaction struct {
	ID          string
	AccountID   int64
	OperationID OperationType
	Amount      float64
	EventDate   time.Time
}
