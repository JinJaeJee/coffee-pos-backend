package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CreatedBy       string
	Name            string
	Picture         string
	Price           float32
	Description     string
	Category        ProductCategory
	IngredientStock []IngredientStock `gorm:"foreignKey:ProductID"`
	OrderDetails    []OrderDetails    `gorm:"foreignKey:ProductID"`
}

type ProductCategory string

const (
	CategoryFood     ProductCategory = "food"
	CategoryDrink    ProductCategory = "drink"
	CategoryDessert  ProductCategory = "dessert"
	CategorySouvenir ProductCategory = "souvenir"
)
