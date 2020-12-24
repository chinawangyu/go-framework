package model

import "time"


type User struct {
	Aid        int       `gorm:"column:aid" json:"aid"`
	UserName   string    `gorm:"column:user_name" json:"user_name"`
	Email      string    `gorm:"column:email" json:"email"`
	UpdateTime time.Time `gorm:"column:email" json:"update_time"`
	CreateTime time.Time `gorm:"-" json:"create_time"`
}
