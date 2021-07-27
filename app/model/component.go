package model

type Component struct {
	Id                  int64   `json:"id"`
	CategoryComponentId *int64  `json:"categoryComponentId"`
	Name                string  `json:"name"`
	Units               string  `json:"units"`
	BuyPrice            float32 `json:"buyPrice"`
	CompanyId           *int64  `json:"companyId"`
	BranchId            *int64  `json:"branchId"`
	Active              bool    `json:"active"`
	CreatedAt           *string `json:"createdAt"`
}

type Components []Component
