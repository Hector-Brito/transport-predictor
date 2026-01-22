package driver

import (
	"context"
	"database/sql"
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


func (r *Repository) GetOne(ctx context.Context, ID int) (*domain.Driver, error){
	query := `SELECT id,first_name,last_name,nickname,created_at,updated_at FROM driver WHERE id = ?;`
	var driver domain.Driver;
	err := r.db.QueryRowContext(ctx, query, ID).Scan(
		&driver.ID,
		&driver.FirstName,
		&driver.LastName,
		&driver.NickName,
		&driver.CreatedAt,
		&driver.UpdatedAt,
	)
	if err != nil {
		return nil,err
	}

	return &driver, nil
}


func (r *Repository) GetAll(ctx context.Context) ([]domain.Driver, error) {
	query := `SELECT id,first_name,last_name,nickname,created_at,updated_at FROM driver;`
	
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	drivers := make([]domain.Driver, 0)

	for rows.Next() {
		var driver domain.Driver

		err := rows.Scan(
			&driver.ID,
			&driver.FirstName,
			&driver.LastName,
			&driver.NickName,
			&driver.CreatedAt,
			&driver.UpdatedAt,
		)
		
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}

	if err = rows.Err();err != nil {
		return nil, err
	}

	return drivers, nil
}

func (r *Repository) Create(ctx context.Context, driver *domain.Driver) (*domain.Driver, error) {
	query := `INSERT INTO driver (first_name, last_name, nickname) VALUES (?, ?, ?);`
	stmt, err := r.db.PrepareContext(ctx, query);

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		driver.FirstName,
		driver.LastName,
		driver.NickName,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	driver.ID = int(id)
	return  driver, nil
}

func (r *Repository) Update(ctx context.Context, ID int, driver *domain.Driver) (*domain.Driver, error) {
	query := `UPDATE driver SET first_name = ?, last_name = ?, nickname = ?, updated_at = ? WHERE id = ?;`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		driver.FirstName,
		driver.LastName,
		driver.NickName,
		driver.UpdatedAt,
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
		return nil, sql.ErrNoRows
	}

	driver.ID = ID
	return  driver, nil
}


func (r *Repository) Delete(ctx context.Context, ID int) error {
	query := `DELETE FROM driver WHERE id = ?;`
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
        return sql.ErrNoRows // El ID no existía
    }

    return nil
	
}