package service

import (
	"gorm.io/gorm"
	"v2/app/model"
)

type Category struct {
	*gorm.DB
}

func NewCategoryService(db *gorm.DB) *Category {
	return &Category{db}
}

func (c *Category) Create(category *model.Category) error {
	if err := c.DB.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) All(scopes ...func(tx *gorm.DB) *gorm.DB) (model.Categories, error) {
	var categories model.Categories

	if err := c.Table("categories").Scopes(scopes...).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
