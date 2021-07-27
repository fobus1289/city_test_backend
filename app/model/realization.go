package model

type Realization struct {
	Id             int64   `json:"id"`
	OrderId        int64   `json:"orderId"`
	Price          float32 `json:"price"`
	ProductName    string  `json:"productName"`
	ReasonCanceled string  `json:"reasonCanceled"`
	WhereRun       string  `json:"whereRun"`
	BranchId       int64   `json:"branchId"`
	ProductId      int64   `json:"productId"`
	Components     string  `json:"components"`
	Status         string  `json:"status"`
	CreatedAt      *string `json:"createdAt"`
}

type Realizations []Realization
