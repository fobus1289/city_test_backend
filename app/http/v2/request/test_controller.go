package request

import (
	"gorm.io/gorm"
	"log"
	"v2/app/service"
)

type Test struct {
	*service.Logger
}

type CategoryComponent struct {
	Id          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	SellPrice   float32 `json:"sell_price,omitempty"`
	Count       float32 `json:"count,omitempty"`
	Units       string  `json:"units,omitempty"`
	Photo       string  `json:"photo,omitempty"`
	Description string  `json:"description,omitempty"`
	Active      bool    `json:"active,omitempty"`
}

type Component struct {
	Id                  int64   `json:"id,omitempty"`
	CategoryComponentId int64   `json:"category_component_id,omitempty"`
	Count               float32 `json:"count,omitempty"`
	Name                string  `json:"name,omitempty"`
	Units               string  `json:"units,omitempty"`
	BuyPrice            float32 `json:"buy_price,omitempty"`
	Active              bool    `json:"active,omitempty"`
}

func (t *Test) CreateProduct() {
	t.Println("CreateProduct")
}

// CreateComponent
// @Security ApiKeyAuth
// @Tags test
// @Description CreateComponent
// @Produce  json
// @Success 200 {object} Component
// @Router /test/create-component [post]
func (t *Test) CreateComponent(c *Component, db *gorm.DB) {
	log.Println(c)
	db.Create(c)
}

// CreateComponentCategory
// @Security ApiKeyAuth
// @Tags test
// @Description CreateComponentCategory
// @Produce  json
// @Success 200 {object} CategoryComponent
// @Router /test/create-component-category [post]
func (t *Test) CreateComponentCategory(cc *CategoryComponent, db *gorm.DB) {
	log.Println(cc)
	db.Create(cc)
}
