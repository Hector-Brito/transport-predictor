package weather

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

func (r *Repository) GetOne(ctx context.Context, ID int) (*domain.Weather, error) {
	query := `SELECT id, name, ground_status, visibility, intensity, temperature, created_at, updated_at FROM weather WHERE id = ?`
	var weather domain.Weather

	err := r.db.QueryRowContext(ctx, query, ID).Scan(
		&weather.ID,
		&weather.Name,
		&weather.GroundStatus,
		&weather.Visibility,
		&weather.Intensity,
		&weather.Temperature,
		&weather.CreatedAt,
		&weather.UpdatedAt,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, err
	}

	return &weather, nil

}

func (r *Repository) GetAll(ctx context.Context) ([]domain.Weather, error) {
	query := `SELECT id, name, ground_status, visibility, intensity, temperature, created_at, updated_at FROM weather`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	weathers := make([]domain.Weather, 0)

	for rows.Next() {
		var weather domain.Weather

		err := rows.Scan(
			&weather.ID,
			&weather.Name,
			&weather.GroundStatus,
			&weather.Visibility,
			&weather.Intensity,
			&weather.Temperature,
			&weather.CreatedAt,
			&weather.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		weathers = append(weathers, weather)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return weathers, nil
}

func (r *Repository) Create(ctx context.Context, weather *domain.Weather) (*domain.Weather, error) {
	query := `INSERT INTO weather (name, ground_status, visibility, intensity, temperature) VALUES (?, ?, ?, ?, ?)`

	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		weather.Name,
		weather.GroundStatus,
		weather.Visibility,
		weather.Intensity,
		weather.Temperature,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	weather.ID = int(id)

	return weather, nil
}

func (r *Repository) Update(ctx context.Context, ID int, weather *domain.Weather) (*domain.Weather, error) {
	query := `UPDATE weather SET name = ?, ground_status = ?, visibility = ?, intensity = ?, temperature = ? WHERE id = ?`

	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		&weather.Name,
		&weather.GroundStatus,
		&weather.Visibility,
		&weather.Intensity,
		&weather.Temperature,
		ID,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, domain.ErrNotFound
	}

	weather.ID = ID

	return weather, nil
}

func (r *Repository) Delete(ctx context.Context, ID int) error {
	query := `DELETE FROM weather WHERE id = ?`

	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		ID,
	)

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}
