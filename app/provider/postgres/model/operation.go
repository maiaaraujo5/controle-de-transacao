package model

import "github.com/uptrace/bun"

type Operation struct {
	bun.BaseModel `bun:"table:operations_types"`

	ID          int64  `bun:"id,pk"`
	description string `bun:"description"`
}
