package domain

import "time"


type TransportLog struct {
	ID int `json:"id"`
	DepartureDate time.Time `json:"departure_date" validate:"required,datetime,ltfield=ArrivalDate"`
	ArrivalDate time.Time `json:"arrival_date" validate:"datetime,gtfield=DepartureDate"`
	Latency time.Time `json:"latency" validate:"datetime"`
	Vehicle Vehicle `json:"vehicle" validate:"required"`
	Weather Weather `json:"weather" validate:"required"`
	DayOfWeek string `json:"day_of_week" validate:"required,oneof=Lunes Martes Miercoles Jueves Viernes Sabado Domingo"`
	IsFortnight bool `json:"is_fortnight" validate:"required" default:"false"`
	Observations string `json:"observations"`
}

type TransportRepository interface {
	GetOne()
	GetAll()
	Create()
	Update()
	Delete()
}