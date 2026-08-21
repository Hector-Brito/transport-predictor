package domain

import (
	"context"
	"time"
)

type TransportLog struct {
	ID            int        `json:"id"`
	DepartureDate *time.Time `json:"departure_date" validate:"required,ltfield=ArrivalDate"`
	ArrivalDate   *time.Time `json:"arrival_date" validate:"gtfield=DepartureDate"`
	Latency       *int       `json:"latency"` //Debe ser calculado automaticamente por el front
	DayOfWeek     *string    `json:"day_of_week" validate:"required,oneof=Monday Tuesday Wednesday Thursday Friday Saturday Sunday"`
	IsFortnight   *bool      `json:"is_fortnight" validate:"required" default:"false"`
	Observations  *string    `json:"observations"`
	WeatherID     *int       `json:"weather_id" validate:"required"`
	VehicleID     *int       `json:"vehicle_id" validate:"required"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type UpdateTransportLogParams struct {
	DepartureDate *time.Time `json:"departure_date" validate:"datetime,ltfield=ArrivalDate"`
	ArrivalDate   *time.Time `json:"arrival_date" validate:"datetime,gtfield=DepartureDate"`
	Latency       *int       `json:"latency"`
	DayOfWeek     *string    `json:"day_of_week" validate:"oneof=Monday Tuesday Wednesday Thursday Friday Saturday Sunday"`
	IsFortnight   *bool      `json:"is_fortnight"`
	Observations  *string    `json:"observations"`
}

// Comenzar por aqui despues de revisar el sql up en migrations
type TransportLogRepository interface {
	GetOne(ctx context.Context, ID int) (*TransportLog, error)
	GetAll(ctx context.Context) ([]TransportLog, error)
	Create(ctx context.Context, transportLog *TransportLog) (*TransportLog, error)
	Update(ctx context.Context, ID int, transportLog *TransportLog) (*TransportLog, error)
	Delete(ctx context.Context, ID int) error
}
