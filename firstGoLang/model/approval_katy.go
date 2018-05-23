package model

import "gcoresys/common/mysql"

type ApprovalKaty struct{
	mysql.BaseModel
	//id
	Id string `gorm:"not null;" json:"id"`
	//name
	Name string `gorm:"not null;" josn:"name"`
	//age
	Age int `gorm:"not null;" json:"age"`
}
