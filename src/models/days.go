package models

import (
	"time"
)

type Days struct{
	DayID   int `gorm:"primary_key;AUTO_INCREMENT"`
	EventID int `gorm "not null"`
	Day     time.Time `gorm "not null"`
}