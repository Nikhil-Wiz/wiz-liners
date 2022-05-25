package services

type Route struct {
	Id int64
	Code string
	VesselName string
	PortOfLoading Port
	PortOfDischarge Port
}