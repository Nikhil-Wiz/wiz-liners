package repositories

import(
	"time"
)

type Cities struct{
	Id int64
	Name string
	Country_id int64
	CreatedAt time.Time
	ModifiedAt time.Time
}