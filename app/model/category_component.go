package model

type CategoryComponent struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	SellPrice   float32 `json:"sellPrice"`
	Count       float32 `json:"count"`
	Units       string  `json:"units"`
	Photo       string  `json:"photo"`
	Description string  `json:"description"`
	CategoryId  *int64  `json:"categoryId"`
	CompanyId   *int64  `json:"companyId"`
	BranchId    *int64  `json:"branchId"`
	Active      bool    `json:"active"`
	CreatedAt   *string `json:"createdAt"`
}

type CategoryComponents []CategoryComponent
