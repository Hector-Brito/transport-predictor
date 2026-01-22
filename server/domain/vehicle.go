package domain

import "time"

type Vehicle struct {
	ID int `json:"id"`
	Name string `json:"name" default:"Standard Bus" validate:"required,oneof='Large Bus' 'Standard Bus' 'Mid-sized Bus' Minibus 'Passenger Pickup' 'Shared Taxi'"`
	NickName string `json:"nickname"`
	Driver Driver `json:"driver" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"min=7,max=8"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VehicleRepository interface {
	GetOne()
	GetAll()
	Create()
	Update()
	Delete()
}