package domain

import (
	"context"
	"time"
)

type Weather struct {
	ID           int       `json:"id"`
	Name         *string   `json:"name" default:"Sunny" validate:"required,oneof=Sunny Foggy Drizzle Storm"`
	GroundStatus *string   `json:"ground_status" default:"Dry" validate:"required,oneof=Dry Wet Puddles/Mud"`
	Visibility   *string   `json:"visibility" default:"Clear" validate:"required,oneof=Clear Foggy"`
	Intensity    *int      `json:"intensity" default:"0" validate:"required,min=0,max=5" `
	Temperature  *int      `json:"temperature" validate:"min=-100,max=100"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateWeatherParams struct {
	Name         *string `json:"name" validate:"required_without_all=Name,omitempty,oneof=Sunny Foggy Drizzle Storm"`
	GroundStatus *string `json:"ground_status" validate:"omitempty,oneof=Dry Wet Puddles/Mud"`
	Visibility   *string `json:"visibility" validate:"omitempty,oneof=Clear Foggy"`
	Intensity    *int    `json:"intensity" validate:"omitempty,min=0,max=5" `
	Temperature  *int    `json:"temperature" validate:"omitempty,min=-100,max=100"`
}

type WeatherRepository interface {
	GetOne(ctx context.Context, ID int) (*Weather, error)
	GetAll(ctx context.Context) ([]Weather, error)
	Create(ctx context.Context, weather *Weather) (*Weather, error)
	Update(ctx context.Context, ID int, weather *Weather) (*Weather, error)
	Delete(ctx context.Context, ID int) error
}
