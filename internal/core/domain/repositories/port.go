package repositories

import "time"

type Port struct {
	Id int64
	Code string
	Name string
	City string
	Country string
	CreatedAt time.Time
	ModifiedAt time.Time
}