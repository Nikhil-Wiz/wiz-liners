package repositories

import (
	"time"
)

type Liners struct {
	Id int64
	Name string
	Code string
	Logo string
	TypeOfLiners string
	CreatedAt time.Time
	ModifiedAt time.Time
}