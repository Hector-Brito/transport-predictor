package transportlog

import (
	"context"
	"database/sql"
	"errors"
	"transport-predictor.com/v2/domain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetOne(ctx context.Context, ID int) (*domain.TransportLog, error) {
	query := `SELECT id, departure_date, arrival_date, latency, day_of_week, is_fortnight, observations, vehicle_id, weather_id, created_at, updated_at FROM transport_log WHERE id = ?`
	var transport_log domain.TransportLog
	err := r.db.QueryRowContext(ctx, query, ID).Scan(
		&transport_log.ID,
		&transport_log.DepartureDate,
		&transport_log.ArrivalDate,
		&transport_log.Latency,
		&transport_log.DayOfWeek,
		&transport_log.IsFortnight,
		&transport_log.Observations,
		&transport_log.VehicleID,
		&transport_log.WeatherID,
		&transport_log.CreatedAt,
		&transport_log.UpdatedAt,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, err
	}

	return &transport_log, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]domain.TransportLog, error) {
	query := `SELECT id, departure_date, arrival_date, latency, day_of_week, is_fortnight, observations, vehicle_id, weather_id, created_at, updated_at FROM transport_log`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	transport_logs := make([]domain.TransportLog, 0)

	for rows.Next() {
		var transport_log domain.TransportLog

		err := rows.Scan(
			&transport_log.ID,
			&transport_log.DepartureDate,
			&transport_log.ArrivalDate,
			&transport_log.Latency,
			&transport_log.DayOfWeek,
			&transport_log.IsFortnight,
			&transport_log.Observations,
			&transport_log.VehicleID,
			&transport_log.WeatherID,
			&transport_log.CreatedAt,
			&transport_log.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		transport_logs = append(transport_logs, transport_log)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transport_logs, nil
}

func (r *Repository) Create(ctx context.Context, transportLog *domain.TransportLog) (*domain.TransportLog, error) {
	query := `INSERT INTO transport_log (departure_date, arrival_date, latency, day_of_week, is_fortnight, observations, vehicle_id, weather_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		transportLog.DepartureDate,
		transportLog.ArrivalDate,
		transportLog.Latency,
		transportLog.DayOfWeek,
		transportLog.IsFortnight,
		transportLog.Observations,
		transportLog.VehicleID,
		transportLog.WeatherID,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	transportLog.ID = int(id)
	return transportLog, nil
}

func (r *Repository) Update(ctx context.Context, ID int, transportLog *domain.TransportLog) (*domain.TransportLog, error) {
	query := `UPDATE transport_log SET departure_date = ?, arrival_date = ?, latency = ?, day_of_week = ?, is_fortnight = ?, observations = ? WHERE id = ?`

	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		&transportLog.DepartureDate,
		&transportLog.ArrivalDate,
		&transportLog.Latency,
		&transportLog.DayOfWeek,
		&transportLog.IsFortnight,
		&transportLog.Observations,
		ID,
	)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, domain.ErrNotFound
	}

	transportLog.ID = ID

	return transportLog, nil
}

func (r *Repository) Delete(ctx context.Context, ID int) error {
	query := `DELETE FROM transport_log WHERE id = ?`

	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}
