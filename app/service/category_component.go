package service

import (
	"gorm.io/gorm"
	"v2/app/model"
)

type CategoryComponent struct {
	*gorm.DB
}

func NewCategoryComponentService(db *gorm.DB) *CategoryComponent {
	return &CategoryComponent{db}
}

func (c *CategoryComponent) Create(component *model.CategoryComponent) error {

	if err := c.Table("category_components").Create(component).Error; err != nil {
		return err
	}

	return nil
}

func (c *CategoryComponent) All() (model.CategoryComponents, error) {
	var categoryComponents model.CategoryComponents

	if err := c.Table("category_components").Find(&categoryComponents).Error; err != nil {
		return nil, err
	}

	return categoryComponents, nil
}
