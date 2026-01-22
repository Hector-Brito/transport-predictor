package domain

import (
	"context"
	"time"
)


type Driver struct {
	ID int `json:"id"`
	FirstName *string `json:"first_name"`
	LastName *string `json:"last_name"`
	NickName string `json:"nickname" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type UpdateDriverParams struct {
	FirstName *string `json:"first_name"`
	LastName *string `json:"last_name"`
	NickName *string `json:"nickname"`
}

type DriverRepository interface {
	GetOne(ctx context.Context, ID int) (*Driver, error)
	GetAll(ctx context.Context) ([]Driver, error)
	Create(ctx context.Context, driver *Driver) (*Driver, error)
	Update(ctx context.Context, ID int, driver *Driver) (*Driver, error)
	Delete(ctx context.Context, ID int) error
}