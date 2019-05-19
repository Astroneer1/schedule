package models


type Event struct{
	EventID   int `gorm:"primary_key;AUTO_INCREMENT"`
	EventName string `gorm "not null"`
	Memo      string
}