package controllers

import (
	"fmt"
	"goshop/app/models"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

// Product contoller which accept request
type Product struct {
	GormController
}

func (p Product) getDb() *gorm.DB {
	return p.Txn
}

//Get product by id
func (p Product) Get() revel.Result {
	db := p.getDb()

	product := models.Product{}
	db.AutoMigrate(&product)
	result := db.Where("id = ?", p.Params.Route.Get("id")).Find(&product)

	if result.Error != nil {
		for _, e := range result.GetErrors() {
			fmt.Println(e)
		}

		return p.RenderJSON(map[string]string{"message": "Error occurs", "error": "true"})
	}

	if result.RecordNotFound() {
		return p.RenderJSON(map[string]string{"message": "Product not found", "id": p.Params.Route.Get("id")})
	}

	category := models.Category{}
	result = db.Model(&product).Related(&category)

	if result.Error == nil && !result.RecordNotFound() {
		product.SetCategory(category)
	}
	return p.RenderJSON(product)
}

// Post  POST params in db
func (p Product) Post() revel.Result {

	db := p.getDb()

	var req map[string]interface{}
	p.Params.BindJSON(&req)
	fmt.Println(req)

	// Find Category

	category := models.Category{}
	result := db.Where("name = ?", req["category"]).First(&category)

	if result.Error != nil {
		for _, e := range result.GetErrors() {
			fmt.Println(e)
		}

		return p.RenderJSON(map[string]string{"message": "Error occurs", "error": "true"})
	}
	if result.RecordNotFound() {
		fmt.Println("Category is not found")
		return p.RenderJSON(map[string]string{"message": "Category not found", "error": "true"})
	}

	fmt.Println(category.ID)

	// Create product
	product := models.Product{}
	product.Name = req["name"].(string)
	product.Price = req["price"].(float64)
	product.CategoryID = category.ID
	product.SetCategory(category)
	db.NewRecord(product)
	db.Create(&product)

	return p.RenderJSON(product)
}
