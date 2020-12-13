package database

import (
  "github.com/kieranroneill/valkyrie/pkg/model"
  "gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
  return db.AutoMigrate(&model.Application{})
}
