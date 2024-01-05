package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	Id   int64  `gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name string `gorm:"not null;unique"`
}

var ErrCategoryAlreadyExists = fmt.Errorf("category already exists")

func CreateCategory(db *gorm.DB, name string) error {
	category := Category{Name: name}

	res := db.Create(&category)
	if res.Error != nil {
		if res.Error.Error() == `ERROR: duplicate key value violates unique constraint "categories_name_key" (SQLSTATE 23505)` {
			return ErrCategoryAlreadyExists
		} else {
			return res.Error
		}

	}

	return nil
}

func GetAvailableCategories(db *gorm.DB) ([]Category, error) {
	var categories []Category

	res := db.Find(&categories)
	if res.Error != nil {
		return nil, res.Error
	}

	return categories, nil
}

func GetCategoryById(db *gorm.DB, id int64) (*Category, error) {
	var category Category

	res := db.First(&category, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &category, nil
}

func EditCategory(db *gorm.DB, id int64, name string) error {
	category := Category{Name: name}

	res := db.Model(&category).Where("id = ?", id).Updates(category)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func DeleteCategoryById(db *gorm.DB, id int64) error {
	res := db.Delete(&Category{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
