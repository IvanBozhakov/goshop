package models

import "gorm.io/gorm"

// Category Entity
type Category struct {
	gorm.Model
	Name string
}
