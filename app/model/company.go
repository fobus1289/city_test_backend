package model

type Company struct {
	Id           int64  `json:"id" db:"id"`
	UserId       *int64 `json:"userId"`
	Name         string  `json:"name" db:"name"`
	INN          int     `json:"inn" db:"inn"`
	FIO          string  `json:"fio" db:"fio"`
	LegalAddress string  `json:"legalAddress" db:"legal_address"`
	Address      string  `json:"address" db:"address"`
	Description  string  `json:"description" db:"description"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	Icon         string  `json:"icon" db:"photo"`
	Site         string  `json:"site" db:"site"`
	Mail         string  `json:"mail" db:"mail"`
	Phone1       string  `json:"phone1" db:"phone1"`
	Phone2       string  `json:"phone2" db:"phone2"`
	Active       bool    `json:"active" db:"active"`
	CreatedAt    *string `json:"createdAt"`
	UpdatedAt    *string `json:"updatedAt"`
	DeletedAt    *string `json:"deletedAt"`
}

type Companies []Company
