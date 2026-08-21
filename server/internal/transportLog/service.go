package transportlog

import (
	"context"
	"time"
	"transport-predictor.com/v2/domain"
)

type Service struct {
	repo domain.TransportLogRepository
}

func NewService(repo domain.TransportLogRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetOne(ctx context.Context, ID int) (*domain.TransportLog, error) {
	return s.repo.GetOne(ctx, ID)
}

func (s *Service) GetAll(ctx context.Context) ([]domain.TransportLog, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, transportLog *domain.TransportLog) (*domain.TransportLog, error) {

	return s.repo.Create(ctx, transportLog)
}

func (s *Service) Update(ctx context.Context, ID int, UpdateTransportLogParams *domain.UpdateTransportLogParams) (*domain.TransportLog, error) {
	transportLog, err := s.repo.GetOne(ctx, ID)

	if err != nil {
		return nil, err
	}

	if UpdateTransportLogParams.DepartureDate != nil {
		transportLog.DepartureDate = UpdateTransportLogParams.DepartureDate
	}

	if UpdateTransportLogParams.ArrivalDate != nil {
		transportLog.ArrivalDate = UpdateTransportLogParams.ArrivalDate
	}

	if UpdateTransportLogParams.Latency != nil {
		transportLog.Latency = UpdateTransportLogParams.Latency
	}

	if UpdateTransportLogParams.DayOfWeek != nil {
		transportLog.DayOfWeek = UpdateTransportLogParams.DayOfWeek
	}

	if UpdateTransportLogParams.IsFortnight != nil {
		transportLog.IsFortnight = UpdateTransportLogParams.IsFortnight
	}

	if UpdateTransportLogParams.Observations != nil {
		transportLog.Observations = UpdateTransportLogParams.Observations
	}

	transportLog.UpdatedAt = time.Now()

	transport_log, err := s.repo.Update(ctx, ID, transportLog)

	if err != nil {
		return nil, err
	}

	return transport_log, nil

}

func (s *Service) Delete(ctx context.Context, ID int) (*domain.TransportLog, error) {
	transport_log, err := s.repo.GetOne(ctx, ID)

	if err != nil {
		return nil, err
	}

	err = s.repo.Delete(ctx, ID)

	if err != nil {
		return nil, err
	}

	return transport_log, nil
}
