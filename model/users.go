package model

type Users struct {
	Id        int    `gorm:"type:int;primary_key"`
	FirstName string `gorm:"type:varchar(255)"`
}
