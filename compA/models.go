package compA

import "gorm.io/gorm"

type CompA struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	name string `json:"name"`
}
