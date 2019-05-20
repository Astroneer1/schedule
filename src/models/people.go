package models


type People struct{
	PeopleID int    `gorm:"primary_key;auto_increment"`
	Name     string `gorm:"not null"`
	Comment  string
	EventID  int    `gorm "not null"`
}