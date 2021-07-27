package model

type Category struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CompanyId *int64 `json:"companyId"`
	Active    bool   `json:"active"`
}

type Categories []Category
