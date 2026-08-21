package vehicle

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

func (r *Repository) GetOne(ctx context.Context, ID int) (*domain.Vehicle, error) {
	query := `SELECT id, name, nickname, license_plate,driver_id, created_at, updated_at FROM vehicle WHERE id = ?`
	var vehicle domain.Vehicle
	err := r.db.QueryRowContext(ctx, query, ID).Scan(
		&vehicle.ID,
		&vehicle.Name,
		&vehicle.NickName,
		&vehicle.LicensePlate,
		&vehicle.DriverID,
		&vehicle.CreatedAt,
		&vehicle.UpdatedAt,
	)

	//manejar error 404 en sql y retornar 404.
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, err
	}

	return &vehicle, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]domain.Vehicle, error) {
	//Defino la consulta a realizar a la base de datos.
	query := `SELECT id, name, nickname, license_plate, driver_id, created_at, updated_at FROM vehicle`
	//Para realizar la consulta utilizo la variable "r" de tipo "*Repository" que interactua con la DB.
	//Se para el contexto y la consulta como parametro.
	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vehicles := make([]domain.Vehicle, 0)

	//Usa rows.Next() para iterar en el resultado de la consulta mediante un bucle for.
	for rows.Next() {
		//Utiliza una variable de tipo Vehicle para almacenar el valor que viene
		var vehicle domain.Vehicle

		// Usa rows.Scan() para enlazar el resultado de la consulta.
		// Utiliza "&" para referenciar la direccion en memoria de cada parametro de la variable.
		// Scan() se encarga de usar la direccion en memoria para enlazar el valor.
		err := rows.Scan(
			&vehicle.ID,
			&vehicle.Name,
			&vehicle.NickName,
			&vehicle.LicensePlate,
			&vehicle.DriverID,
			&vehicle.CreatedAt,
			&vehicle.UpdatedAt,
		)

		//En caso de que ocurra un error, lo retorna.
		if err != nil {
			return nil, err
		}

		//Agrega a la slice de vehicles un vehicle.
		vehicles = append(vehicles, vehicle)
	}
	//Si rows tiene algun error, lo guarda en la variable ya creada "err" y luego evalua la condicion de si es distinto a nil, retorna el error.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return vehicles, nil

}

func (r *Repository) Create(ctx context.Context, vehicle *domain.Vehicle) (*domain.Vehicle, error) {
	query := `INSERT INTO vehicle (name, nickname, license_plate, driver_id) VALUES (?, ?, ?, ?)`
	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		vehicle.Name,
		vehicle.NickName,
		vehicle.LicensePlate,
		vehicle.DriverID,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	vehicle.ID = int(id)

	return vehicle, nil

}

func (r *Repository) Update(ctx context.Context, ID int, vehicle *domain.Vehicle) (*domain.Vehicle, error) {
	query := `UPDATE vehicle SET name = ?, nickname = ?, license_plate = ? WHERE id = ?`

	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	//Falta cambio de driver_id - en el sentido de que se puede cambiar el driver.
	//para luego
	result, err := stmt.ExecContext(
		ctx,
		&vehicle.Name,
		&vehicle.NickName,
		&vehicle.LicensePlate,
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

	vehicle.ID = ID

	return vehicle, nil

}

func (r *Repository) Delete(ctx context.Context, ID int) error {
	query := `DELETE FROM vehicle WHERE id = ?`

	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(
		ctx,
		ID,
	)

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
