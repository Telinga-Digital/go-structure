package models

import "time"

type User struct {
	ID        int       `gorm:"primary_key;type:bigint(20) unsigned autoIncrement"`
	Email     string    `gorm:"column:email;type:varchar(255);unique"`
	Password  string    `gorm:"column:password;type:varchar(255)"`
	ApiKey    string    `gorm:"column:api_key;type:varchar(255);unique"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp"`
}
