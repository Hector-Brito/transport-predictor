package domain

import (
	"context"
	"time"
)

// Es necesario crear un struct que refleje la db y otro para la respuesta.
type Vehicle struct {
	ID int `json:"id"`
	//usar puntero en name para que pueda ser nulo y coloque el default.
	Name     *string `json:"name" default:"Standard Bus" validate:"required,oneof='Large Bus' 'Standard Bus' 'Mid-sized Bus' Minibus 'Passenger Pickup' 'Shared Taxi'"`
	NickName *string `json:"nickname"`
	//Plantear hacer uso de puntero en campo DriverID
	DriverID     int       `json:"driver_id" validate:"required"`
	LicensePlate *string   `json:"license_plate" validate:"min=7,max=8"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Usar en la creacion de la feature Update.
type UpdateVehicleParams struct {
	Name         *string `json:"name,omitempty" validate:"omitempty,oneof='Large Bus' 'Standard Bus' 'Mid-sized Bus' Minibus 'Passenger Pickup' 'Shared Taxi'"`
	NickName     *string `json:"nickname,omitempty" validate:"required_without_all=Name,omitempty"`
	LicensePlate *string `json:"license_plate" validate:"omitempty,min=7,max=8"`
}

type VehicleRepository interface {
	GetOne(ctx context.Context, ID int) (*Vehicle, error)
	GetAll(ctx context.Context) ([]Vehicle, error)
	Create(ctx context.Context, vehicle *Vehicle) (*Vehicle, error)
	Update(ctx context.Context, ID int, vehicle *Vehicle) (*Vehicle, error)
	Delete(ctx context.Context, ID int) error
}
