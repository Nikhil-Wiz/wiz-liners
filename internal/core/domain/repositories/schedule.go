package repositories

import (
	"time"
	"google.golang.org/genproto/googleapis/type/datetime"
)

type Schedules struct {
	Id int64
	IsAvailable bool
	LinerId int64
	TransitDays int64
	FreeDays int64
	PortOfLoading int64
	PortOfDischarge int64
	PortOfLoadingDate datetime.DateTime
	PortOfDischargeDate datetime.DateTime
	CreatedAt time.Time
	ModifiedAt time.Time
}