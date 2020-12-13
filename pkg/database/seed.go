package database

import (
  "encoding/json"
  "fmt"
  "github.com/kieranroneill/valkyrie/pkg/model"
  "gorm.io/gorm"
  "io/ioutil"
  "os"
  "path"
  "strings"
)

func getDataDir() (string, error) {
  pwd, err := os.Getwd()
  if err != nil {
    return "", err
  }

  return fmt.Sprintf("%s/data/", pwd), nil
}

func containsApplication(data []*model.Application, alias string) bool {
  for _, m := range data {
    if m.Alias == alias {
      return true
    }
  }

  return false
}

func seedApplications(db *gorm.DB, file []byte) error {
  var jsondata []*model.Application
  var dbdata []*model.Application

  // Add the json data to the interface.
  if err := json.Unmarshal(file, &jsondata); err != nil {
    return err
  }

  if err := db.Find(&dbdata).Error; err != nil {
    return err
  }

  for _, d := range jsondata {
    if containsApplication(dbdata, d.Alias) {
      if err := db.Where("alias = ?", d.Alias).Updates(d).Error; err != nil {
        return err
      }

      continue
    }

    if err := db.Create(d).Error; err != nil {
      return err
    }
  }

  return nil
}

func RunSeeds(db *gorm.DB) error {
  dir, err := getDataDir()
  if err != nil {
    return err
  }

  // Get all the json files from the data directory.
  files, err := ioutil.ReadDir(dir)
  if err != nil {
    return err
  }

  // Seed each collection.
  for _, f := range files {
    clnName := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
    pthNme := fmt.Sprintf("%s%s", dir, f.Name())
    file, err := ioutil.ReadFile(pthNme)
    if err != nil {
      return err
    }

    switch clnName {
    case "applications":
      if err = seedApplications(db, file); err != nil {
        return err
      }
      break
    default:
      break
    }
  }

  return nil
}
