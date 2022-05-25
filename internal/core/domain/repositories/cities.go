package repositories

import(
	"time"
)

type Cities struct{
	Id int64
	Name string
	Country_id int64
	Created_At time.Time
	Modified_At time.Time
}