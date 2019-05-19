package models

type Days struct{
	DayID   int `gorm:"primary_key;AUTO_INCREMENT"`
	EventID int `gorm "not null"`
	Day     string `gorm "not null"`
}