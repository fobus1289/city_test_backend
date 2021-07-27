package model

type Branch struct {
	Id           int64   `json:"id"`
	CompanyId    *int64  `json:"companyId"`
	UserId       *int64  `json:"userId"`
	Name         string  `json:"name"`
	Address      string  `json:"address"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	Active       bool    `json:"active"`
	Phone1       string  `json:"phone1"`
	Phone2       string  `json:"phone2"`
	LegalAddress string  `json:"legalAddress"`
	Photo        string  `json:"photo"`
	UntilDate    *string `json:"untilDate"`
	CreatedAt    *string `json:"createdAt"`
	UpdatedAt    *string `json:"updatedAt"`
	DeletedAt    *string `json:"deletedAt"`
}

type Branches []Branch

func (b *Branch) FixUtilDate() {
	if b.UntilDate != nil {
		date := *b.UntilDate
		date = date[:10]
		b.UntilDate = &date
	}
}
