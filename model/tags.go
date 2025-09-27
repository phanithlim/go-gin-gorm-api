package model

type Tags struct {
	Id   int    `grom:"type:int;primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}
