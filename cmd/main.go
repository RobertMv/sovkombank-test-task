package main

import (
	"clinic"
	"clinic/internal/handler"
	"clinic/internal/repository"
	"clinic/internal/service"
	"database/sql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("Error occured during config initializing: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(viper.GetString("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %s", err.Error())
	}

	migrateDB(db)

	repos := repository.NewRepositories(db)
	services := service.NewServices(repos, db)
	handlers := handler.NewHandlers(services)

	srv := new(clinic.Server)
	if err := srv.Run(viper.GetString("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server")
	}
}

func initConfig() error {
	viper.SetConfigFile("сonfig/.env")
	return viper.ReadInConfig()
}

func migrateDB(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Couldn't get database instance for running migrations; %s", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "clinic", driver)
	if err != nil {
		log.Fatalf("Couldn't create migrate instance; %s", err.Error())
	}

	if err := m.Up(); err != nil {
		log.Printf("Couldn't run database migrations; %s", err.Error())
	} else {
		log.Println("Database migration was run successfully")
	}
}
