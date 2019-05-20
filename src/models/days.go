package models

type Days struct{
	DayID   int    `gorm:"primary_key;auto_increment"`
	EventID int    `gorm "not null"`
	Day     string `gorm "not null"`
}
