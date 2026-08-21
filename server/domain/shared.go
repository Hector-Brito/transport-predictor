package domain

import (
	"errors"
)

// Sentinel errors pertenecen a la parte de la logica de negocio.
var (
	ErrNotFound = errors.New("resource not found")
)
