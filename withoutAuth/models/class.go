package models

type (
	Class struct {
		ID       uint      `gorm:"primary_key;autoIncrement:true"`
		Name     string    `json:"name"`
		Capacity uint      `json:"capacity"`
		Students []Student `json:"-"`
	}
)
