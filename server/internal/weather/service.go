package weather

import (
	"transport-predictor.com/v2/domain"
)

type Service struct {
	repo *domain.WeatherRepository
}