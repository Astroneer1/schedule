package models


type Event struct{
	EventID   int    `gorm:"primary_key;auto_increment"`
	EventName string `gorm "not null"`
	Memo      string
}
