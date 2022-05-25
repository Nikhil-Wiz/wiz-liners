package repositories

import "time"

type Ports struct {
	Code string
	Name string
	Type string
	City_Id int64
	State string
	Latitude float64
	Longitude float64
	CreatedAt time.Time
	ModifiedAt time.Time
}