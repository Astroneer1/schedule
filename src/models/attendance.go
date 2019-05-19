package models


type Attendamce struct{
	AttendamceID int `gorm:"primary_key;AUTO_INCREMENT"`
	Attendamce   int `gorm "not null"`
	DayID        int `gorm "not null"`
	PeopleID     int `gorm "not null"`
}