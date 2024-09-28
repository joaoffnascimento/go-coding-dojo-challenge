package repository

import (
	"encoding/json"
	"log"

	"gorm.io/gorm"

	"go-challenge/src/model"
)

func CreateProduct[P any](db *gorm.DB, product *P) (*P, error) {
	if err := db.Create(product).Error; err != nil {
		return nil, err
	}

	// Log the created product as JSON
	productJSON, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	log.Printf("Created product: %s", productJSON)

	return product, nil
}

func GetProductByID(db *gorm.DB, id uint) (*model.Product, error) {
	var product model.Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func UpdateProductByID(db *gorm.DB, id uint, updatedData map[string]interface{}) (*model.Product, error) {
	var product model.Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&product).Updates(updatedData).Error; err != nil {
		return nil, err
	}

	// Log the updated product as JSON
	productJSON, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	log.Printf("Updated product: %s", productJSON)

	return &product, nil
}
