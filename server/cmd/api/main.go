package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"transport-predictor.com/v2/internal/database"
	"transport-predictor.com/v2/internal/driver"
	"transport-predictor.com/v2/internal/server"
)



func main() {
	err := godotenv.Load()
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

	driverRepository := driver.NewRepository(db);
	driverService := driver.NewService(driverRepository);
	driverHandler := driver.NewHandler(driverService)

	srv := server.NewServer();
	handlers := &server.Handlers{
		Driver:driverHandler,
	}
	srv.RegisterRoutes(handlers)
	log.Printf("Starting server on %v",backend_port)
	if err := srv.Run(backend_port);err != nil{
		log.Fatal(err)
	}
}