package model

type Application struct {
  Alias string `json:"alias"`
  Description string `json:"description"`
  ID int `json:"id" gorm:"primarykey"`
  Name string `json:"name"`
  Public bool `json:"public"`
  Url string `json:"url"`
}
