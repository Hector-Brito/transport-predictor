# Variables por defecto (se pueden sobreescribir al ejecutar make)
DB_URL=sqlite3://db.db
MIGRATIONS_PATH=migrations 
# Sube todas las migraciones o N pasos si se especifica (ej: make migrate-up N=-1) 
migrate-up: 
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up $(N)
# Revierte todas las migraciones o N pasos si se especifica (ej: make migrate-down N=1)
migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down $(N)

# Crea una nueva migración (ej: make migrate-create name=create_users_table)
migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

# Fuerza una versión específica (útil si la DB queda en estado "dirty" tras un error)
migrate-force:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $(V)

# Muestra el numero de migraciones dentro de la base de datos.
show-migrations:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version
