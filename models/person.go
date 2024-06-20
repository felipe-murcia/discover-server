package models

import (
	"time"
)

type Person struct {
	ID        int32     `gorm:"primaryKey;column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Address   string    `gorm:"column:address" json:"address"`
	Phone     uint      `gorm:"column:phone" json:"phone"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli" json:"updated_at"`
}

func (Person) TableName() string {
	return "persons"
}
