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
	Created_At time.Time
	Modified_At time.Time
}