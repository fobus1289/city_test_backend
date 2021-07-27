package service

import (
	"gorm.io/gorm"
	"v2/app/model"
)

type Component struct {
	*gorm.DB
}

func NewComponentService(db *gorm.DB) *Component {
	return &Component{db}
}

func (c *Component) All() (model.Components, error) {

	var components model.Components

	if err := c.Table("components").Find(&components).Error; err != nil {
		return nil, err
	}

	return components, nil
}

func (c *Component) Create(component *model.Component) error {
	if err := c.Table("components").Create(component).Error; err != nil {
		return err
	}

	return nil
}
