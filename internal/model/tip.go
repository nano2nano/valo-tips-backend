package model

import (
	"gorm.io/gorm"
)

type Tip struct {
	gorm.Model
	Title        string `gorm:"not null"`
	StandImgPath string `gorm:"column:stand_img_path;not null"`
	AimImgPath   string `gorm:"column:aim_img_path;not null"`
	Description  string `gorm:"not null"`
	SideID       uint   `gorm:"not null"`
	Side         Side
	MapUUID      string `gorm:"column:map_uuid;not null"`
	AgentUUID    string `gorm:"column:agent_uuid;not null"`
	AbilitySlot  string `gorm:"not null"`
	Good         int    `gorm:"not null;default:0"`
	Bad          int    `gorm:"not null;default:0"`
}

type Tips []Tip

func (t *Tips) GetAll(tx *gorm.DB) error {
	return tx.Preload("Side").Preload("Map").Preload("Ability").Find(t).Order("ID asc").Error
}

func (t *Tip) Create(tx *gorm.DB) error {
	return tx.Create(t).Error
}

func (t *Tip) Get(tx *gorm.DB, id uint) error {
	return tx.Take(t, id).Error
}

func (t *Tip) Delete(tx *gorm.DB, id uint) error {
	return tx.Delete(t, id).Error
}
