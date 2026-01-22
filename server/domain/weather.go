package domain

import "time"


type Weather struct {
	ID int `json:"id"`
	Name string `json:"name" default:"Sunny" validate:"required,oneof=Sunny Foggy Drizzle Storm"`
	GroundStatus string `json:"ground_status" default:"Dry" validate:"required,oneof=Dry Wet Puddles/Mud"`
	Visibility string `json:"visibility" default:"Clear" validate:"required,oneof=Clear Foggy"` 
	Intensity int `json:"intensity" default:"0" validate:"required,min=0,max=5" `
	Temperature int `json:"temperature" validate:"min=-100,max=100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type WeatherRepository interface {
	GetOne()
	GetAll()
	Create()
	Update()
	Delete()
}