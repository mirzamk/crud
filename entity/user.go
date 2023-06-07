package entity

import "time"

type User struct {
	//gorm.Model
	Id        uint   `gorm:"primary_key"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
