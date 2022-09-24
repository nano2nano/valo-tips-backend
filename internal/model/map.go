package model

import "gorm.io/gorm"

type Map struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type Maps []Map

func (m *Map) Save(tx *gorm.DB) error {
	return tx.Create(m).Error
}

func (m *Map) Get(tx *gorm.DB, id uint) error {
	return tx.Take(m, id).Error
}

func (m *Map) Delete(tx *gorm.DB, id uint) error {
	return tx.Delete(m, id).Error
}

func (m *Maps) GetAll(tx *gorm.DB) error {
	return tx.Find(m).Order("ID asc").Error
}
