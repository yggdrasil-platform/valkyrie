package application

import (
  "github.com/kieranroneill/valkyrie/pkg/config"
  "github.com/kieranroneill/valkyrie/pkg/database"
  "github.com/kieranroneill/valkyrie/pkg/logger"
  "gorm.io/gorm"
)

type Application struct {
	Config *config.Config
  Database *gorm.DB
}

func New() (*Application, error) {
  // Connect to the DB
  db, err := database.New()
  if err != nil {
    return nil, err
  }

  // Run db migrations
  if err = database.RunMigrations(db); err != nil {
    logger.Error.Printf("failed to run database migrations: %s", err)
  }

  return &Application{
    Config: config.New(),
    Database: db,
  }, nil
}
