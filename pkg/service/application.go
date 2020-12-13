package service

import (
  "github.com/kieranroneill/valkyrie/pkg/logger"
  "github.com/kieranroneill/valkyrie/pkg/model"
  "gorm.io/gorm"
)

type ApplicationService struct {
	database *gorm.DB
}

func(s *ApplicationService) Get() []*model.Application {
  var ms []*model.Application

  result := s.database.Find(&ms)
  if result.Error != nil {
    logger.Error.Printf(result.Error.Error())
    return ms
  }

  return ms
}

func(s *ApplicationService) GetById(id int) *model.Application {
	var m model.Application

	result := s.database.First(&m, id)
	if result.Error != nil {
		logger.Error.Printf(result.Error.Error())
		return nil
	}

	return &m
}

func(s *ApplicationService) GetByAlias(alias string) *model.Application {
 var m model.Application

 result := s.database.Where("alias = ?", alias).First(&m)
 if result.Error != nil {
   logger.Error.Printf(result.Error.Error())
   return nil
 }

 return &m
}

func NewApplicationService(db *gorm.DB) *ApplicationService {
	return &ApplicationService{database: db}
}
