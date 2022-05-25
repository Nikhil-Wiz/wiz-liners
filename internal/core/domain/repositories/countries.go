package repositories

import(
	"time"
)

type Countries struct{
	Id int64
	Name string
	Iso_code string
	Currency_code string
	Created_At time.Time
	Modified_At time.Time
}