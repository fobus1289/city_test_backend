package model

type Product struct {
	Id                int64             `json:"id"`
	Name              string            `json:"name"`
	Photo             string            `json:"photo"`
	Price             float32           `json:"price"`
	Description       string            `json:"description"`
	WhereRun          string            `json:"whereRun"`
	CompanyId         *int64            `json:"companyId"`
	BranchId          *int64            `json:"branchId"`
	Active            bool              `json:"active"`
	CreatedAt         *string           `json:"createdAt"`
	ProductComponents ProductComponents `json:"components"`
}

type Products []Product

type ProductComponent struct {
	Id                 int64 `json:"id"`
	ProductId          int64 `json:"productId"`
	ComponentId        int64 `json:"componentId"`
	CanAdd             bool  `json:"canAdd"`
	CanRemove          bool  `json:"canRemove"`
	ChangedPriceAdd    bool  `json:"changedPriceAdd"`
	ChangedPriceRemove bool  `json:"changedPriceRemove"`
}

type ProductComponents []ProductComponent

func (p ProductComponents) SetProductId(id int64) {
	for i := 0; i < len(p); i++ {
		p[i].ProductId = id
	}
}
