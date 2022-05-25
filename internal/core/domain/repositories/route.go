package repositories

import "time"

type Route struct {
	Id int64
	Code string
	VesselName string
	PortOfLoading int64
	PortOfDischarge int64
	CreatedAt time.Time
	ModifiedAt time.Time
}