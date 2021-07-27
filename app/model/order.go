package model

type Order struct {
	Id             int64        `json:"id"`
	Price          float32      `json:"price"`
	Status         string       `json:"status"`
	Number         int          `json:"number"`
	BranchId       int64        `json:"branchId"`
	Description    string       `json:"description"`
	ReasonCanceled string       `json:"reasonCanceled"`
	CustomerId     int64        `json:"customerId"`
	WorkerId       int64        `json:"workerId"`
	Realizations   Realizations `json:"realizations" gorm:"foreignKey:order_id"`
	CreatedAt      *string      `json:"createdAt"`
}

type Orders []Order

type COrder struct {
	Description string `json:"description"`
	Products    Order  `json:"order"`
}

type DescriptionOrder struct {
	Description string   `json:"description"`
	Products    Products `json:"products"`
}
