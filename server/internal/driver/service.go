package driver

import (
	"context"
	"time"

	"transport-predictor.com/v2/domain"
)


type Service struct {
	repo domain.DriverRepository
}

func NewService(repo domain.DriverRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetOne(ctx context.Context, ID int) (*domain.Driver, error){
	return s.repo.GetOne(ctx, ID)
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Driver, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, driver *domain.Driver) (*domain.Driver, error) {
	return s.repo.Create(ctx, driver)
}

func (s *Service) Update(ctx context.Context, ID int, UpdateDriver *domain.UpdateDriverParams) (*domain.Driver, error) {
	driver, err := s.repo.GetOne(ctx, ID)
	if err != nil {
		return nil, err
	}

	if UpdateDriver.FirstName != nil {
		driver.FirstName = UpdateDriver.FirstName
	}
	
	if UpdateDriver.LastName != nil {
		driver.LastName = UpdateDriver.LastName
	}
	
	if UpdateDriver.NickName != nil {
		driver.NickName = *UpdateDriver.NickName
	}
	
	driver.UpdatedAt = time.Now()

	driver, err = s.repo.Update(ctx, ID, driver)

	if err != nil {
		return nil, err
	}
	return  driver, nil
}

func (s *Service) Delete(ctx context.Context, ID int) (*domain.Driver, error) {
	driver, err := s.repo.GetOne(ctx, ID)

	if err != nil {
		return nil, err
	}

	err = s.repo.Delete(ctx, ID)
    if err != nil {
        return nil, err
    }

	return  driver, nil
}