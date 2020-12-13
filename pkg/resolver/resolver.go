package resolver

import (
  "github.com/kieranroneill/valkyrie/pkg/config"
  "gorm.io/gorm"
)

type Resolver struct {
  Config *config.Config
  Database *gorm.DB
}
