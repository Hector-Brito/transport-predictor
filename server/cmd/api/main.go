package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"transport-predictor.com/v2/internal/database"
	"transport-predictor.com/v2/internal/driver"
	"transport-predictor.com/v2/internal/server"
	transportlog "transport-predictor.com/v2/internal/transportLog"
	"transport-predictor.com/v2/internal/vehicle"
	"transport-predictor.com/v2/internal/weather"
)

func main() {
	err := godotenv.Load("/home/hectorbrito/Documents/Programming/transport-predictor/server/.env")
	if err != nil {
		log.Fatal(err)
		return
	}
	db_filepath := os.Getenv("DB_FILEPATH")
	backend_port := os.Getenv("BACKEND_PORT")

	if db_filepath == "" {
		log.Fatal("Set your 'DB_FILEPATH' environment variable.")
		return
	}

	db, err := database.NewSQLiteConnection(db_filepath)

	if err != nil {
		log.Fatal("Cannot connect to SQLite database.")
		return
	}
	defer db.Close()

	driverRepository := driver.NewRepository(db)
	vehicleRepository := vehicle.NewRepository(db)
	weatherRepository := weather.NewRepository(db)
	transportLogRepository := transportlog.NewRepository(db)

	driverService := driver.NewService(driverRepository)
	vehicleService := vehicle.NewService(vehicleRepository)
	weatherService := weather.NewService(weatherRepository)
	transportLogService := transportlog.NewService(transportLogRepository)

	driverHandler := driver.NewHandler(driverService)
	vehicleHandler := vehicle.NewHandler(vehicleService)
	weatherHandler := weather.NewHandler(weatherService)
	transportLogHandler := transportlog.NewHandler(transportLogService)

	srv := server.NewServer()
	handlers := &server.Handlers{
		Driver:      driverHandler,
		Vehicle:     vehicleHandler,
		Weather:     weatherHandler,
		TranportLog: transportLogHandler,
	}

	srv.RegisterRoutes(handlers)
	log.Printf("Starting server on %v", backend_port)
	if err := srv.Run(backend_port); err != nil {
		log.Fatal(err)
	}
}
