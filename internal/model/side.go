package model

import "gorm.io/gorm"

type Side struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"not null"`
}

type Sides []Side

func (s *Side) Create(tx *gorm.DB) error {
	return tx.Create(s).Error
}

func (s *Sides) GetAll(tx *gorm.DB) error {
	return tx.Find(s).Order("ID asc").Error
}
