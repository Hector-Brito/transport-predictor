package vehicle

import (
	"context"
	"time"
	"transport-predictor.com/v2/domain"
)

type Service struct {
	repo domain.VehicleRepository
}

func NewService(repo domain.VehicleRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetOne(ctx context.Context, ID int) (*domain.Vehicle, error) {
	return s.repo.GetOne(ctx, ID)
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Vehicle, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, vehicle *domain.Vehicle) (*domain.Vehicle, error) {
	return s.repo.Create(ctx, vehicle)
}

func (s *Service) Update(ctx context.Context, ID int, UpdateVehicle *domain.UpdateVehicleParams) (*domain.Vehicle, error) {
	vehicle, err := s.repo.GetOne(ctx, ID)

	if err != nil {
		return nil, err
	}

	if UpdateVehicle.Name != nil {
		vehicle.Name = UpdateVehicle.Name
	}

	if UpdateVehicle.NickName != nil {
		vehicle.NickName = UpdateVehicle.NickName
	}

	if UpdateVehicle.LicensePlate != nil {
		vehicle.LicensePlate = UpdateVehicle.LicensePlate
	}

	vehicle.UpdatedAt = time.Now() //Esto no hace nada, hay que arreglarlo.
	//Se supone que este atributo se debe actualizar en la db.
	//Sin embargo, solo se actualiza en la variable, pero luego que se completa todo, no se guarda.
	//Porque no esta implementado siquiera en repository para su actualizacion.
	//Es posible arreglar y seguir haciendo uso de esta logica, o implementar triggers.
	//Los triggers acomplarian logica de negocio en la db (segun mi entendimiento).

	vehicle, err = s.repo.Update(ctx, ID, vehicle)

	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (s *Service) Delete(ctx context.Context, ID int) (*domain.Vehicle, error) {
	vehicle, err := s.repo.GetOne(ctx, ID)

	if err != nil {
		return nil, err
	}

	err = s.repo.Delete(ctx, ID)

	if err != nil {
		return nil, err
	}

	return vehicle, nil
}
