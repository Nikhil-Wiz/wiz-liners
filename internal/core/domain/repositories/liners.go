package repositories

import (
	"time"
)

type Liners struct {
	Id int64
	Name string
	Code string
	Type string  //sea or air
	Logo string	
	CreatedAt time.Time
	ModifiedAt time.Time
}