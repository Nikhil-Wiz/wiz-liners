package services

type Schedules struct {
	Id int64
	IsAvailable bool
	LinerId int64
	TransitDays int64
	FreeDays int64
	PortOfLoading int64
	PortOfDischarge int64
	PortOfLoadingDate Port
	PortOfDischargeDate Port
}