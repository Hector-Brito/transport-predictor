package weather

import (
	"context"
	"time"

	"transport-predictor.com/v2/domain"
)

type Service struct {
	repo domain.WeatherRepository
}

func NewService(repo domain.WeatherRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetOne(ctx context.Context, ID int) (*domain.Weather, error) {
	return s.repo.GetOne(ctx, ID)
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Weather, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, weather *domain.Weather) (*domain.Weather, error) {
	return s.repo.Create(ctx, weather)

}

func (s *Service) Update(ctx context.Context, ID int, updateWeather *domain.UpdateWeatherParams) (*domain.Weather, error) {
	weather, err := s.repo.GetOne(ctx, ID)

	if err != nil {
		return nil, err
	}

	if updateWeather.Name != nil {
		weather.Name = updateWeather.Name
	}

	if updateWeather.GroundStatus != nil {
		weather.GroundStatus = updateWeather.GroundStatus
	}

	if updateWeather.Visibility != nil {
		weather.Visibility = updateWeather.Visibility
	}

	if updateWeather.Intensity != nil {
		weather.Intensity = updateWeather.Intensity
	}
	if updateWeather.Temperature != nil {
		weather.Temperature = updateWeather.Temperature
	}

	weather.UpdatedAt = time.Now()

	weather, err = s.repo.Update(ctx, ID, weather)

	if err != nil {
		return nil, err
	}
	return weather, nil

}

func (s *Service) Delete(ctx context.Context, ID int) (*domain.Weather, error) {
	weather, err := s.repo.GetOne(ctx, ID)

	if err != nil {
		return nil, err
	}

	err = s.repo.Delete(ctx, ID)

	if err != nil {
		return nil, err
	}

	return weather, nil
}
