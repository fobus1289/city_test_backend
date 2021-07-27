package service

import "gorm.io/gorm"

type Realization struct {
	*gorm.DB
}

func NewRealizationService(db *gorm.DB) *Realization {
	return &Realization{db}
}
