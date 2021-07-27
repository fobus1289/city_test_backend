package service

import (
	"gorm.io/gorm"
	"v2/app/model"
)

type Order struct {
	*gorm.DB
}

var number = 0

func NewOrderService(db *gorm.DB) *Order {
	return &Order{db}
}

func (o *Order) All() (model.Orders, error) {
	var orders model.Orders

	if err := o.Table("orders").Where("status in (?)", "created", "inProcessing").
		Preload("Realizations").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *Order) ForClients() (model.Orders, error) {
	var orders model.Orders

	if err := o.Table("orders").
		Where("status != ?", "canceled").
		Where("status != ?", "delivered").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *Order) StatusUpdate(order *model.Order) (*model.Order, error) {

	var (
		tx  = o.Begin()
		err error
	)

	defer commit(tx, func() error {
		return err
	})

	if err = tx.Table("orders").Where("id=?", order.Id).
		Update("status", order.Status).Update("reason_canceled", order.ReasonCanceled).
		Error; err != nil {
		return nil, err
	}

	var realizationIds []int64

	for _, realization := range order.Realizations {
		realizationIds = append(realizationIds, realization.Id)
	}

	if err = tx.Table("realizations").Where("id in (?)", realizationIds).
		Update("status", order.Status).Update("reason_canceled", order.ReasonCanceled).
		Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) Create(descriptionOrder *model.DescriptionOrder, jwtUser *model.JwtUser) (*model.Order, error) {

	var (
		products     = descriptionOrder.Products
		realizations model.Realizations
		prices       float32 = 0
		tx                   = o.Begin()
		err          error
	)

	defer commit(tx, func() error {
		return err
	})

	for _, product := range products {
		prices += product.Price
	}

	number++

	var order = &model.Order{
		Price:       prices,
		Status:      "created",
		Number:      number,
		Description: descriptionOrder.Description,
		BranchId:    1,
		CustomerId:  0,
		WorkerId:    jwtUser.Id,
	}

	if err = tx.Table("orders").Omit("CreatedAt").Create(&order).Error; err != nil {
		return nil, err
	}

	for _, product := range products {
		//var productComponents model.ProductComponents
		//
		//if err = tx.Table("product_components").Where("product_id", product.Id).Find(&productComponents).Error; err != nil {
		//	return nil, err
		//}

		realizations = append(realizations, model.Realization{
			OrderId:     order.Id,
			Price:       product.Price,
			ProductName: product.Name,
			WhereRun:    product.WhereRun,
			BranchId:    1,
			ProductId:   product.Id,
			Status:      "created",
		})
	}

	if err = tx.Table("realizations").Omit("CreatedAt").Create(&realizations).Error; err != nil {
		return nil, err
	}

	order.Realizations = realizations

	return order, err
}
