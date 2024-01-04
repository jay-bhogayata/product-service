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
