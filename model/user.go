package model

import (
	"time"
)

type User struct {
	UserID        	uint      	`json:"user_id" gorm:"primary_key;column:user_id;not null"`
	AuthID			string		`json:"auth_id" gorm:"column:auth_id;default:null"`
	UserName      	string    	`json:"user_name" gorm:"column:user_name;default:null"`
	Email      		string    	`json:"email" gorm:"column:email;default:null"`
	Birth      		string    	`json:"birth" gorm:"column:birth;default:null"`
	Sex      		string    	`json:"sex" gorm:"column:sex;default:null"`
	Phone      		string    	`json:"phone" gorm:"column:phone;default:null"`
	UserStatus      string    	`json:"user_status" gorm:"column:user_status;not null;default:1"`
	Admin      		string    	`json:"admin" gorm:"column:admin;not null;default:0"`
	CreatedTime 	time.Time 	`json:"created_time" gorm:"column:created_time;not null;default:CURRENT_TIMESTAMP"`
	UpdatedTime 	time.Time 	`json:"updated_time" gorm:"column:updated_time;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (User) TableName() string {
	return "User"
}