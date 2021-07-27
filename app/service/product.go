package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"v2/app/model"
)

type Product struct {
	*gorm.DB
}

func NewProductService(db *gorm.DB) *Product {
	return &Product{db}
}

func (p *Product) All() (model.Products, error) {

	var products model.Products

	if err := p.Table("products").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p *Product) Create(product *model.Product) error {

	var (
		err               error
		tx                = p.Begin()
		productComponents = product.ProductComponents
	)
	log.Println(productComponents)
	defer commit(tx, func() error {
		return err
	})

	if err = tx.Table("products").Omit("ProductComponents").
		Create(&product).Error; err != nil {
		return err
	}

	if len(productComponents) > 0 {
		productComponents.SetProductId(product.Id)
		log.Println(productComponents)
		if err = tx.Table("product_components").Clauses(clause.Insert{Modifier: "IGNORE"}).
			Create(&productComponents).Error; err != nil {
			return err
		}
	}

	return nil
}
