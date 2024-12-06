package models

type (
	Student struct {
		ID      uint   `gorm:"primary_key;autoIncrement:true"`
		Name    string `json:"name"`
		Age     int    `json:"age"`
		ClassID uint   `json:"class_id"`
	}
)
