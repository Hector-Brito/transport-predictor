package transportlog

import "github.com/go-playground/validator/v10"

type Handler struct {
	service *Service
	validate *validator.Validate
}