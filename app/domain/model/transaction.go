package model

import "time"

type Transaction struct {
	ID        int64
	AccountID int64
	Operation OperationType
	Amount    float64
	EventDate time.Time
}
