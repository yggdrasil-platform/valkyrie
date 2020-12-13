package model

type Application struct {
  Alias string `json:"alias"`
  Description string `json:"description"`
  ID int `json:"id" gorm:"primarykey"`
  Internal bool `json:"internal"`
  Name string `json:"name"`
  Url string `json:"url"`
}
