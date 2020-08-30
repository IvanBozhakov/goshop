package models

import "gorm.io/gorm"

// Product Entity
type Product struct {
	gorm.Model
	Name       string
	Price      float64
	CategoryID uint
	category   Category `gorm:"foreignkey:CategoryID"`
}

// SetCategory add category for this product
func (p *Product) SetCategory(category Category) {
	p.category = category
}
