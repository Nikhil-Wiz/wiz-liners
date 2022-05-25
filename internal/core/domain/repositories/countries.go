package repositories

import(
	"time"
)

type Countries struct{
	Id int64
	Name string
	Iso_code string
	Currency_code string
	CreatedAt time.Time
	ModifiedAt time.Time
}